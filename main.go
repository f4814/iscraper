package main

import (
	"sync"

	"github.com/ahmdrz/goinsta/v2"
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to load config file: ", err)
	}

	coreConfig := NewCoreConfig(viper.GetStringMap("core"))
	databaseConfig := NewDatabaseConfig(viper.GetStringMap("arangodb"))
	instagramConfig := NewInstagramConfig(viper.GetStringMap("instagram"))
	scraperConfig := NewScraperConfig(viper.GetStringMap("scraper"))

	// Setup Logging
	switch coreConfig.LogLevel {
	case "TRACE":
		log.SetLevel(log.TraceLevel)
		log.SetReportCaller(true)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	}

	// Connect to ArangoDB
	var (
		dbConfig   driver.ClientConfig
		httpConfig http.ConnectionConfig
	)

	log.Info(databaseConfig.Authentication)
	if databaseConfig.Authentication != nil {
		username := databaseConfig.Authentication["username"]
		password := databaseConfig.Authentication["password"]
		dbConfig.Authentication = driver.BasicAuthentication(username, password)
	}

	httpConfig.Endpoints = databaseConfig.URL

	conn, err := http.NewConnection(httpConfig)
	dbConfig.Connection = conn
	if err != nil {
		log.Fatal(err)
	}

	client, err := driver.NewClient(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	log.WithFields(log.Fields{
		"url": databaseConfig.URL,
	}).Info("Connected to ArangoDB")

	// Initalize database
	helper := NewDBHelper(client, databaseConfig.Database)

	// Initialize goinsta
	insta := goinsta.New(instagramConfig.Username, instagramConfig.Password)

	if instagramConfig.Proxy != "" {
		insta.SetProxy(instagramConfig.Proxy, true)
	}

	if err := insta.Login(); err != nil {
		log.Fatal(err)
	} else {
		log.WithFields(log.Fields{
			"username": viper.GetString("instaUser"),
		}).Info("Authenticated at Instagram")
	}

	// Start Scrapers
	var wg sync.WaitGroup
	queue := make(chan goinsta.User, 100)

	for i := 0; i < scraperConfig.Scrapers; i++ {
		config := scraperConfig
		config.ID = 0

		go Scraper(config, queue, helper, wg)
		wg.Add(1)

		log.WithFields(log.Fields{
			"id": i,
		}).Debug("Started Scraper")
	}

	// Start queue
	// go Queue(queue, helper, wg)
	// wg.Add(1)
	// log.Debug("Started Queue")

	// Root User
	rootUsername := scraperConfig.RootUser
	if rootUsername != "" {
		root, err := insta.Profiles.ByName(rootUsername)
		if err != nil {
			log.Warn(err)
		} else {
			log.WithFields(log.Fields{
				"username": rootUsername,
			}).Info("Loaded root user")
			queue <- *root
		}
	} else {
		log.Debug("No root user specified. Conitinuing with queue from db")
	}

	log.Info("Startup Finished")
	wg.Wait()
}
