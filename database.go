package main

import (
	"context"

	"github.com/ahmdrz/goinsta"
	"github.com/f4814/iscraper/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

func saveUser(d *mongo.Database, u *goinsta.User) {
	// Prepare models.User struct
	var data models.User
	u.Sync()

	data.FromIG(u)
	followers := u.Followers()
	following := u.Following()
	for followers.Next() {
		for _, v := range followers.Users {
			data.Followers = append(data.Followers, v.ID)
		}
	}
	for following.Next() {
		for _, v := range following.Users {
			data.Following = append(data.Following, v.ID)
		}
	}

	// Convert to BSON and save to Database
	b, err := bson.Marshal(data)
	if err != nil {
		panic(err)
	}

	if _, err := d.Collection("users").InsertOne(context.Background(), b); err != nil {
		panic(err)
	}

	log.Info("Inserted User: ", v.Username)
}
