package main

import (
	"context"
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
)

func main() {
	log.SetLevel(log.DebugLevel)
	kingpin.Parse()

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

	// Get Root user
	insta := goinsta.New(*username, *password)
	if err := insta.Login(); err != nil {
		panic(err)
	}

	user, err := insta.Profiles.ByName("sas_weingarten")

	if err != nil {
		panic(err)
	}

	user.Sync()
	following := user.Following()

	for following.Next() {
		for _, u := range following.Users {
			saveUser(dat, &u)

		}
	}
}
