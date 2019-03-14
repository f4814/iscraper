package main

import (
	"fmt"
	"time"

	"sync"

	"github.com/ahmdrz/goinsta/v2"
	"github.com/f4814/iscraper/models"
	log "github.com/sirupsen/logrus"
)

func Scraper(config ScraperConfig, queue chan goinsta.User, helper DBHelper,
	wg *sync.WaitGroup) {

	defer wg.Done()
	var (
		modelUser *models.User
		logFields log.Fields
	)

	scrapeFeedMedia := func(media *goinsta.FeedMedia,
		relation func(*models.User, *models.Item)) {

		if err := media.Sync(); err != nil {
			// log.WithFields(logFields).Warn(err)
		}

		for media.Next() {
			if err := media.Error(); err != nil {
				log.WithFields(logFields).Warn(err)
			}

			for _, item := range media.Items {
				if err := item.SyncLikers(); err != nil {
					log.WithFields(logFields).Warn(err)
				}

				itemModel := models.NewItem(item)
				itemModel.ScrapedAt = time.Now()
				helper.SaveItem(itemModel)
				relation(modelUser, itemModel)

				// TODO: Comments

				if config.Scrape["item_likers"] {
					for _, liker := range item.Likers {
						likerModel := models.NewUser(liker)
						helper.SaveUser(likerModel)
						helper.UserLikes(likerModel, itemModel)
					}
				}

				fmt.Println("%#v", item.Tags)
				if config.Scrape["item_tags"] {
					for i, t := range models.NewTags(item.FbUserTags) {
						in := item.FbUserTags.In[i]
						t.FBUserTag = true

						m := models.NewUser(in.User)
						helper.SaveUser(m)

						helper.ItemTags(itemModel, m, t)
					}

					for _, v := range item.Tags.In {
						for i, t := range models.NewTags(v) {
							in := v.In[i]
							t.FBUserTag = false

							m := models.NewUser(in.User)
							helper.SaveUser(m)

							helper.ItemTags(itemModel, m, t)
						}
					}
				}
			}
		}
	}

	scrapeUsers := func(users *goinsta.Users,
		relation func(*models.User, *models.User)) {

		for users.Next() {
			if err := users.Error(); err != nil {
				log.WithFields(logFields).Warn(err)
			}

			for _, other := range users.Users {
				modelOther := models.NewUser(other)
				helper.SaveUser(modelOther)
				relation(modelUser, modelOther)
			}
		}
	}

	for {
		user, more := <-queue

		if !more {
			log.Info("Queue Closed. Stopping Scraper")
			return
		}

		logFields = log.Fields{
			"username": user.Username,
			"scraper":  config.ID,
		}

		log.WithFields(logFields).Info("Scraping User")

		if err := user.Sync(); err != nil {
			log.WithFields(logFields).Warn(err)
		}

		modelUser = models.NewUser(user)
		modelUser.ScrapedAt = time.Now()
		helper.SaveUser(modelUser)

		// Scrape Followers
		if config.Scrape["user_followers"] {
			log.WithFields(logFields).Debug("Scraping Followers")
			followers := user.Followers()
			scrapeUsers(followers, helper.UserFollowed)
		}

		// Scrape following
		if config.Scrape["user_following"] {
			log.WithFields(logFields).Debug("Scraping following")
			following := user.Following()
			scrapeUsers(following, helper.UserFollows)
		}

		// Scrape user feed
		if config.Scrape["user_feed"] {
			log.WithFields(logFields).Debug("Scraping Feed")
			feed := user.Feed()
			scrapeFeedMedia(feed, helper.UserPosts)
		}

		// Scrape Media user is Tagged in
		if config.Scrape["user_tagged"] {
			log.WithFields(logFields).Debug("Scraping Media tagging user")
			if tags, err := user.Tags(nil); err != nil {
				log.WithFields(logFields).Warn(err)
			} else {
				scrapeFeedMedia(tags, helper.UserTagged)
			}
		}

		// TODO Stories, Highlights
	}
}
