package main

import (
	"sync"

	"github.com/ahmdrz/goinsta"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

func Scrape(wg *sync.WaitGroup, c chan *goinsta.User, d *mongo.Database) {
	defer wg.Done()

	for {
		user, more := <-c

		if !more {
			log.Debug("Stopping Worker")
			return
		}

		saveUser(c, d, user)
	}
}
