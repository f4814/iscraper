package main

import (
	"context"
	"sync"
	"time"

	"github.com/ahmdrz/goinsta"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose  = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	username = kingpin.Flag("user", "Instagram username.").Required().String()
	password = kingpin.Flag("password", "Instagram password.").Required().String()
	mongodb  = kingpin.Flag("mongodb", "Mongodb connection.").Required().String()
	database = kingpin.Flag("database", "database").Required().String()
	workers  = kingpin.Flag("workers", "Number of Workers").Required().Int()
	root     = kingpin.Flag("root", "Root user").Required().String()
)

func main() {
	kingpin.Parse()

	if *verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// Connect to mongodb
	client, err := mongo.NewClient(*mongodb)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	dat := client.Database(*database)
	log.Info("Connected to mongodb on ", *mongodb)

	// Get Root user
	insta := goinsta.New(*username, *password)
	if err := insta.Login(); err != nil {
		panic(err)
	}
	log.Info("Authenticated as ", *username)

	user, err := insta.Profiles.ByName("lexodexo.de")
	if err != nil {
		panic(err)
	}
	log.Info("Loaded root user ", user.Username)

	// Initialize workers
	users := make(chan *goinsta.User, 500000)
	var wg sync.WaitGroup
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		log.Debug("Starting worker (", i, ")")
		go Scrape(&wg, users, dat)
	}
	users <- user
	wg.Wait()

}
