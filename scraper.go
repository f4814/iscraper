package main

import (
	"time"

	"sync"

	"github.com/ahmdrz/goinsta/v2"
	"github.com/f4814/iscraper/models"
	log "github.com/sirupsen/logrus"
)

func Scraper(config ScraperConfig, queue chan goinsta.User, helper DBHelper,
	wg sync.WaitGroup) {

	defer wg.Done()

	for {
		user, more := <-queue

		if !more {
			log.Info("Queue Closed. Stopping Scraper")
			return
		}

		log.WithFields(log.Fields{
			"username": user.Username,
			"scraper":  config.ID,
		}).Info("Scraping User")

		if err := user.Sync(); err != nil {
			log.WithFields(log.Fields{
				"username": user.Username,
				"scraper":  config.ID,
			}).Warn(err)
		}

		modelUser := models.NewUser(user)
		modelUser.ScrapedAt = time.Now()
		helper.SaveUser(modelUser)

		// Scrape Followers
		log.WithFields(log.Fields{
			"username": user.Username,
			"scraper":  config.ID,
		}).Debug("Scraping Followers")
		followers := user.Followers()
		scrapeUsers(modelUser, followers, helper, helper.UserFollowed)

		// Scrape following
		log.WithFields(log.Fields{
			"username": user.Username,
			"scraper":  config.ID,
		}).Debug("Scraping following")
		following := user.Following()
		scrapeUsers(modelUser, following, helper, helper.UserFollows)

		// Scrape user feed
		log.WithFields(log.Fields{
			"username": user.Username,
			"scraper":  config.ID,
		}).Debug("Scraping Feed")
		feed := user.Feed()
		scrapeFeedMedia(modelUser, feed, helper, helper.UserPosts)

		// Scrape Media user is Tagged in
		log.WithFields(log.Fields{
			"username": user.Username,
			"scraper":  config.ID,
		}).Debug("Scraping Media tagging user")
		if tags, err := user.Tags(nil); err != nil {
			log.WithFields(log.Fields{
				"username": user.Username,
				"scraper":  config.ID,
			}).Warn(err)
		} else {
			scrapeFeedMedia(modelUser, tags, helper, helper.UserTagged)
		}

		// TODO Stories, Highlights
	}
}

func scrapeFeedMedia(user *models.User, media *goinsta.FeedMedia,
	helper DBHelper, relation func(*models.User, *models.Item)) {

	if err := media.Sync(); err != nil {
		log.Warn(err)
	}

	for media.Next() {
		if err := media.Error(); err != nil {
			log.Warn(err)
		}

		for _, item := range media.Items {
			if err := item.SyncLikers(); err != nil {
				log.Warn(err)
			}

			itemModel := models.NewItem(item)
			helper.SaveItem(itemModel)
			relation(user, itemModel)

			// TODO Item tags
			// TODO: Comments

			for _, liker := range item.Likers {
				likerModel := models.NewUser(liker)
				helper.SaveUser(likerModel)
				helper.UserLikes(likerModel, itemModel)
			}
		}
	}
}

func scrapeUsers(user *models.User, users *goinsta.Users, helper DBHelper,
	relation func(*models.User, *models.User)) {

	for users.Next() {
		if err := users.Error(); err != nil {
			log.Warn(err)
		}

		for _, other := range users.Users {
			modelOther := models.NewUser(other)
			helper.SaveUser(modelOther)
			relation(user, modelOther)
		}
	}
}
