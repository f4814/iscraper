package main

import (
	"sync"

	"github.com/ahmdrz/goinsta"
	"github.com/f4814/iscraper/models"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

func Scrape(wg *sync.WaitGroup, d *mongo.Database, queue chan goinsta.User) {
	defer wg.Done()

	for {
		user, more := <-queue

		if !more {
			log.Debug("Stopping Worker")
			return
		}

		if checkUser(d, user.Username) {
			log.Debug("Not rescraping user: ", user.Username, " (", len(queue), ")")
			continue
		}

		scraped := scrapeUser(&user)
		feed := user.Feed()
		saveUser(d, scraped)

		log.Info("Scraping user items: ", scraped.Username)
		for feed.Next() {
			for _, i := range feed.Items {
				item := scrapeItem(&i)
				saveItem(d, item)
			}
		}

		toQueue := append(scraped.FollowingStructs, scraped.FollowerStructs...)

		for _, u := range toQueue {
			if u.ID != user.ID && !checkUser(d, u.Username) {
				queue <- u
				log.Debug("Queueing user: ", u.Username, " (", len(queue), ")")
			} else {
				log.Debug("Not requeueing user: ", u.Username)
			}
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
			data.FollowerStructs = append(data.FollowerStructs, v)
		}
	}

	for following.Next() {
		for _, v := range following.Users {
			data.Following = append(data.Following, v.ID)
			data.FollowingStructs = append(data.FollowingStructs, v)
		}
	}

	return &data
}

// Scrape a goinsta.Item into a models.User
func scrapeItem(item *goinsta.Item) *models.Item {
	var data models.Item

	item.SyncLikers()
	data.FromIG(item)

	return &data
}
