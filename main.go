package main

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/ahmdrz/goinsta"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose    = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	username   = kingpin.Flag("user", "Instagram username.").Required().String()
	password   = kingpin.Flag("password", "Instagram password.").Required().String()
	mongodb    = kingpin.Flag("mongodb", "Mongodb connection.").Required().String()
	database   = kingpin.Flag("database", "database").Required().String()
	workers    = kingpin.Flag("workers", "Number of Workers").Required().Int()
	root       = kingpin.Flag("root", "Root user").Required().String()
	debug      = kingpin.Flag("debug", "Show log caller method").Bool()
	cooldown   = kingpin.Flag("cooldown", "Cooldown on api block").Required().String()
	bufferSize = kingpin.Flag("bufferSize", "Buffer Size").Required().Int()
)

func main() {
	// Parse CLI Flags
	kingpin.Parse()

	if *debug {
		log.SetReportCaller(true)
	}

	if *verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// Connect to mongodb
	client, err := mongo.NewClient(*mongodb)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
		os.Exit(1)
	}

	dat := client.Database(*database)
	log.Info("Connected to mongodb on ", *mongodb)

	// Authenticate
	insta := goinsta.New(*username, *password)
	if err := insta.Login(); err != nil {
		log.Fatal("Failed to authenticate with instagram: ", err)
		os.Exit(1)
	}
	log.Info("Authenticated as ", *username)

	// Get root user
	rootUser, err := insta.Profiles.ByName(*root)
	if err != nil {
		log.Fatal("Failed to find root user: ", err)
	}
	log.Info("Loaded root user: ", rootUser.Username)

	// Parse cooldown
	cooldownParsed, err := time.ParseDuration(*cooldown)
	if err != nil {
		log.Fatal("Unable to parse cooldown time: ", err)
		os.Exit(1)
	}

	// Initialize workers
	queue := make(chan goinsta.User, *bufferSize)
	var wg sync.WaitGroup

	for i := 0; i < *workers; i++ {
		wg.Add(1)
		log.Debug("Starting worker (", i, ")")
		go Scrape(&wg, dat, queue, cooldownParsed)
	}

	queue <- *rootUser

	go func() {
		for {
			QueueNext(queue, dat, insta, cooldownParsed)
		}
	}()

	wg.Wait()
}
