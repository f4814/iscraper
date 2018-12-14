package main

import (
	"context"

	"github.com/f4814/iscraper/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

// Save a models.User
func saveUser(d *mongo.Database, user *models.User) error {
	if err := writeBSON(d, "users", user); err != nil {
		// Do not return duplicate key errors
		switch t := err.(type) {
		default:
			return err
		case mongo.WriteErrors:
			for _, e := range t {
				if e.Code != 11000 {
					return err
				}
			}
			log.Warn("Failed to add duplicate user: ", user.Username)
		}
	}

	log.Debug("Added User: ", user.Username)

	return nil
}

// Save a models.Item
func saveItem(d *mongo.Database, item *models.Item) error {
	if err := writeBSON(d, "items", item); err != nil {
		// Do not return duplicate key errors
		switch t := err.(type) {
		default:
			return err
		case mongo.WriteErrors:
			for _, e := range t {
				if e.Code != 11000 {
					return err
				}
			}
			log.Warn("Failed to add duplicate item")
		}
	}

	log.Debug("Added Item")

	return nil
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
		return err
	}

	return nil
}
