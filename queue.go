package main

import (
	"context"
	"time"

	"github.com/ahmdrz/goinsta"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	log "github.com/sirupsen/logrus"
)

type Queued struct {
	ID   int64        `bson:"_id"`
	User goinsta.User `bson:"user"`
	At   time.Time    `bson:"at"`
}

// Send the user to the queue channel. If the channel is full, put it into the database
func QueueOne(curr goinsta.User, user goinsta.User, queue chan goinsta.User,
	d *mongo.Database) error {

	if curr.ID == user.ID || checkUser(d, user.Username) {
		log.Debug("Not requeueing user: ", user.Username)
		return nil
	}

	select {
	case queue <- user:
		log.Debug("Queueing user: ", user.Username, " (", len(queue), ")")
	default:
		log.Debug("Queueing user to database: ", user.Username)

		var q Queued
		q.ID = user.ID
		q.At = time.Now()
		q.User = user

		b, err := bson.Marshal(q)
		if err != nil {
			log.Warn(err)
			return err
		}

		_, err = d.Collection("queue").InsertOne(context.Background(), b)

		if err != nil && !isDuplicateError(err) {
			return err
		}
	}

	return nil
}

// Execute QueueOne on each element of a slice
func QueueMany(curr goinsta.User, users []goinsta.User, queue chan goinsta.User,
	d *mongo.Database) error {

	for _, v := range users {
		if err := QueueOne(curr, v, queue, d); err != nil { // TODO Use Insert Many
			return err
		}
	}

	return nil
}

// Load queued users from the database and send them into the queue channel
func QueueNext(queue chan goinsta.User, d *mongo.Database,
	insta *goinsta.Instagram, cooldown time.Duration) {

	collection := d.Collection("queue")
	ctx := context.Background()
	findOpts := options.Find().SetSort(bsonx.Doc{{"at", bsonx.Int32(1)}})
	deleteOpts := options.Delete()

	cur, err := collection.Find(ctx, nil, findOpts)
	defer cur.Close(ctx)

	if err != nil {
		log.Fatal(err)
	}

	var (
		q Queued
		u goinsta.User
	)

	for cur.Next(ctx) {
		if err := cur.Decode(&q); err != nil {
			log.Fatal(err)
		}

		collection.DeleteOne(ctx, bsonx.Doc{{"_id", bsonx.Int64(q.ID)}}, deleteOpts)

		u = q.User
		u.SetInstagram(insta)

		if !checkUser(d, u.Username) {
			log.Debug("Queuing user from database: ", u.Username)
			queue <- u
		} else {
			log.Debug("Not requeueing user from database: ", u.Username)
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
