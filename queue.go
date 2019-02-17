package main

import (
	"sync"

	"github.com/ahmdrz/goinsta/v2"
)

func Queue(queue chan goinsta.User, helper DBHelper, wg sync.WaitGroup) {
	defer wg.Done()

	for {
	}
}
