package main

import (
	"context"

	"github.com/ahmdrz/goinsta"
	"github.com/f4814/iscraper/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

func saveUser(ch chan *goinsta.User, d *mongo.Database, u *goinsta.User) {
	if checkUser(d, u.Username) {
		log.Info("User ", u.Username, " already in database. Skipping")
		return
	}

	// Prepare models.User struct
	var data models.User
	u.Sync()

	data.FromIG(u)
	followers := u.Followers()
	following := u.Following()

	for followers.Next() {
		for _, v := range followers.Users {
			data.Followers = append(data.Followers, v.ID)
			if !checkUser(d, v.Username) {
				ch <- &v
				log.Debug("Queueing user: ", v.Username)
			}
		}
	}

	for following.Next() {
		for _, v := range following.Users {
			data.Following = append(data.Following, v.ID)
			if !checkUser(d, v.Username) {
				ch <- &v
				log.Debug("Queueing user: ", v.Username)
			}
		}
	}

	if err := writeBSON(d, "users", data); err != nil {
		panic(err)
	}

	log.Info("Added User: ", u.Username)
}

// Checks whether the given user already exists
func checkUser(d *mongo.Database, username string) bool {
	filter := bson.D{{"username", username}}
	res := d.Collection("users").FindOne(context.Background(), filter)

	var u models.User
	err := res.Decode(&u)

	if err != nil {
		return false
	}

	if u.Username == username {
		return true
	}

	return false
}

// Writes a BSON struct to a collection
func writeBSON(d *mongo.Database, collection string, val interface{}) error {
	b, err := bson.Marshal(val)
	if err != nil {
		log.Warn(err)
		return err
	}

	_, err = d.Collection(collection).InsertOne(context.Background(), b)

	if err != nil {
		log.Warn(err)
		return err
	}

	return nil
}
