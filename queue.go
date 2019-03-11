package main

import (
	"sync"

	"github.com/ahmdrz/goinsta/v2"
	"github.com/f4814/iscraper/models"
	log "github.com/sirupsen/logrus"
)

func Queue(queue chan goinsta.User, helper DBHelper, insta *goinsta.Instagram,
	wg sync.WaitGroup) {

	defer wg.Done()

	for {
		query := `FOR u IN users
					FILTER u.scraped_at == "1970-01-01T01:00:00+01:00"
					SORT u.added_at ASC
					RETURN u`
		cur, err := helper.DB.Query(nil, query, nil)
		defer cur.Close()

		if err != nil {
			log.Warn(err)
		}

		for cur.HasMore() {
			var newModelUser models.User
			_, err := cur.ReadDocument(nil, &newModelUser)

			if err != nil {
				log.Warn(err)
				continue
			}

			newUser := models.NewGoinstaUser(&newModelUser)
			newUser.SetInstagram(insta)
			queue <- *newUser
			log.WithFields(log.Fields{
				"username": newUser.Username,
			}).Debug("Queued user")
		}
	}
}
