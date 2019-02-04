package main

import (
	"time"

	"github.com/ahmdrz/goinsta"
	"github.com/f4814/iscraper/models"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

func Scrape(d *mongo.Database, queue chan goinsta.User, exit chan bool,
	cooldown time.Duration) {

	for {
		select {
		case _ = <-exit:
			return
		default:
		}

		user, more := <-queue

		if !more {
			return
		}

		if checkUser(d, user.Username) {
			log.Debug("Not rescraping user: ", user.Username, " (", len(queue), ")")
			continue
		}

		scraped := scrapeUser(&user, cooldown)
		feed := user.Feed()
		saveUser(d, *scraped)

		log.Info("Scraping user items: ", scraped.Username)
		for feed.Next() {
			for _, i := range feed.Items {
				item := scrapeItem(&i, cooldown)
				saveItem(d, item)
			}
		}

		toQueue := append(scraped.FollowingStructs, scraped.FollowerStructs...)

		if err := QueueMany(user, toQueue, queue, d); err != nil {
			log.Fatal("Failed to Queue user: ", err)
		}
	}
}

// Scrape a goinsta.User into a models.User
func scrapeUser(user *goinsta.User, cooldown time.Duration) *models.User {
	var data models.User

	log.Info("Scraping user: ", user.Username)

	for {
		if err := user.Sync(); err == nil {
			break
		}

		log.Warn("Blocked by API, cooling down worker for ", cooldown, " seconds")
		time.Sleep(cooldown)
	}

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
func scrapeItem(item *goinsta.Item, cooldown time.Duration) *models.Item {
	var data models.Item

	for {
		if err := item.SyncLikers(); err == nil {
			break
		}

		log.Warn("Blocked by API, cooling down worker for ", cooldown, " seconds")
		time.Sleep(cooldown)
	}

	item.Comments.Sync()

	data.FromIG(item)

	return &data
}
