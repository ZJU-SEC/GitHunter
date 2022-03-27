package util

import (
	"GitHunter/config"
	"fmt"
	"github.com/gocolly/colly"
	"math/rand"
	"sync"
	"time"
)

func RandomString() string {
	const bytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(b)
}

// CommonCollector return a base collector
func CommonCollector() *colly.Collector {
	c := colly.NewCollector()

	// set random `User-Agent`
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		r.Headers.Set("Authorization", "token "+requestGitHubToken())
	})

	c.OnError(func(r *colly.Response, err error) {
		if r.StatusCode == 404 {
			return
		}
		if config.DEBUG {
			fmt.Println("[debug]", r.StatusCode, r.Request.URL)
		}
		retryRequest(r.Request, config.TRYOUT)
	})

	return c
}

func retryRequest(r *colly.Request, maxRetries int) int {
	retriesLeft := maxRetries
	if x, ok := r.Ctx.GetAny("retriesLeft").(int); ok {
		retriesLeft = x
	}
	if retriesLeft > 0 {
		r.Ctx.Put("retriesLeft", retriesLeft-1)
		time.Sleep(time.Duration(config.TRYOUT-retriesLeft) * time.Second)
		r.Retry()
	} else {
		fmt.Println("! cannot fetch", r.URL)
	}
	return retriesLeft
}

var count = 0

// requestGitHubToken return one token from the list and loop the pointer
func requestGitHubToken() string {
	l := len(config.GITHUB_TOKEN)
	var mutex sync.Mutex
	mutex.Lock()

	token := config.GITHUB_TOKEN[count%l] // select one token
	count++                               // auto increment

	if count > l { // prevent overflow
		count -= l
	}

	mutex.Unlock()
	return token
}
