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
	viper.SetDefault("databaseURL", "http://localhost:8529")
	viper.SetDefault("database", "iscraper")
	viper.SetDefault("authentication", false)
	viper.SetDefault("debug", false)
	viper.SetDefault("scrapers", 3)

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to load config file: ", err)
	}

	// Setup Logging
	if viper.GetBool("debug") {
		log.SetLevel(log.TraceLevel)
		log.SetReportCaller(false)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// Connect to ArangoDB
	var (
		dbConfig   driver.ClientConfig
		httpConfig http.ConnectionConfig
	)

	if viper.GetBool("authentication") {
		username := viper.GetString("username")
		password := viper.GetString("password")
		dbConfig.Authentication = driver.BasicAuthentication(username, password)
	}

	httpConfig.Endpoints = []string{viper.GetString("databaseURL")}

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
		"url": viper.GetString("databaseURL"),
	}).Info("Connected to ArangoDB")

	// Initalize database
	dbName := viper.GetString("database")
	helper := NewDBHelper(client, dbName)

	// Initialize goinsta
	insta := goinsta.New(viper.GetString("instaUser"), viper.GetString("instaPW"))
	insta.SetProxy("127.0.0.1:3128", true)
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

	for i := 0; i < viper.GetInt("scrapers"); i++ {
		go Scraper(queue, helper, wg)
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
	rootUsername := viper.GetString("rootUser")
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
