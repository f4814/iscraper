package main

import (
	"sync"

	"github.com/ahmdrz/goinsta"
	"github.com/f4814/iscraper/models"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

func Scrape(wg *sync.WaitGroup, d *mongo.Database, queue chan *goinsta.User) {
	defer wg.Done()

	for {
		user, more := <-queue

		if !more {
			log.Debug("Stopping Worker")
			return
		}

		if checkUser(d, user.Username) {
			log.Debug("Not rescraping user: ", user.Username)
			continue
		}

		scraped := scrapeUser(user);
		saveUser(d, scraped)

		for _, u := range scraped.FollowingStructs {
			queue <- u
		}
		for _, u := range scraped.FollowerStructs {
			queue <- u
		}
	}
}

// Scrape a goinsta.User into a models.User
func scrapeUser(user *goinsta.User) *models.User {
	var data models.User
	user.Sync()

	log.Info("Scraping user: ", user.Username)

	data.FromIG(user)
	followers := user.Followers()
	following := user.Following()

	for followers.Next() {
		for _, v := range followers.Users {
			data.Followers = append(data.Followers, v.ID)
			data.FollowerStructs = append(data.FollowerStructs, &v)
		}
	}

	for following.Next() {
		for _, v := range following.Users {
			data.Following = append(data.Following, v.ID)
			data.FollowingStructs = append(data.FollowingStructs, &v)
		}
	}

	return &data
}
