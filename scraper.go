package main

import (
	"time"

	"sync"

	"github.com/ahmdrz/goinsta/v2"
	"github.com/f4814/iscraper/models"
	log "github.com/sirupsen/logrus"
)

func Scraper(queue chan goinsta.User, helper DBHelper, wg sync.WaitGroup) {
	defer wg.Done()

	for {
		user, more := <-queue

		if !more {
			log.Info("Queue Closed. Stopping Scraper")
			return
		}

		log.WithFields(log.Fields{
			"username": user.Username,
		}).Info("Scraping User")

		if err := user.Sync(); err != nil {
			log.Warn(err)
		}

		modelUser := models.NewUser(user)
		modelUser.ScrapedAt = time.Now()
		helper.SaveUser(modelUser)

		// Scrape Followers
		log.WithFields(log.Fields{
			"username": user.Username,
		}).Debug("Scraping Followers")
		followers := user.Followers()
		for followers.Next() {
			if err := followers.Error(); err != nil {
				log.Warn(err)
			}

			for _, follower := range followers.Users {
				modelFollower := models.NewUser(follower)
				helper.SaveUser(modelFollower)
				helper.UserFollowed(modelUser, modelFollower)
			}
		}

		// Scrape following
		log.WithFields(log.Fields{
			"username": user.Username,
		}).Debug("Scraping following")
		following := user.Following()
		for following.Next() {
			if err := following.Error(); err != nil {
				log.Warn(err)
			}

			for _, follows := range following.Users {
				modelFollows := models.NewUser(follows)
				helper.SaveUser(modelFollows)
				helper.UserFollows(modelUser, modelFollows)
			}
		}

		// Scrape user feed
		// TODO: Comments
		log.WithFields(log.Fields{
			"username": user.Username,
		}).Debug("Scraping Feed")
		feed := user.Feed()
		for feed.Next() {
			if err := feed.Error(); err != nil {
				log.Warn(err)
			}

			for _, item := range feed.Items {
				if err := item.SyncLikers(); err != nil {
					log.Warn(err)
				}

				itemModel := models.NewItem(item)
				helper.SaveItem(itemModel)
				helper.UserPosts(modelUser, itemModel)

				for _, liker := range item.Likers {
					likerModel := models.NewUser(liker)
					helper.SaveUser(likerModel)
					helper.UserLikes(likerModel, itemModel)
				}
			}
		}

		// TODO: Stories
	}
}
