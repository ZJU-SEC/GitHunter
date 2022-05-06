package web

import (
	"GitHunter/config"
	"GitHunter/model"
	"GitHunter/util"
	"encoding/json"
	"fmt"

	// "math"

	"github.com/gocolly/colly"
	"github.com/shomali11/parallelizer"
)

// Resp is a simple api response handler struct
type Resp struct {
	Count int          `json:"total_count"`
	Repos []model.Repo `json:"items"`
}

var MAX_STAR = 10000

func Crawl() {
	crawlWithOption(MAX_STAR, 1<<23)
	crawlWithOption(config.MIN_STAR, MAX_STAR)
}

// For the cursor, this function return the repos having star in range [min, max]
func crawlWithOption(min, max int) {
	if max < min { return }
	
	var resp Resp
	// send one request, if bigger than 1000, divide them into two parts
	c := util.CommonCollector()
	c.OnResponse(func(r *colly.Response) {
		json.Unmarshal(r.Body, &resp)
	})

	queryURL := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s+stars:%d..%d",
		config.LANGUAGE, min, max)

	if config.DEBUG {
		fmt.Println("visiting", queryURL)
	}

	c.Visit(queryURL)

	if resp.Count == 0 {
		return
	} else if resp.Count > 1000 && min < max {
		mid := (min + max) / 2
		// Making sure the granularity is 1
		crawlWithOption(mid+1, max)
		crawlWithOption(min, mid)
		return
	}

	model.CreateRepoBatch(resp.Repos)
	pages := (resp.Count-1)/30 + 1
	if pages > 100 {
		page = 100
	}

	// iterate from page 2 to the final page
	group := parallelizer.NewGroup(
		parallelizer.WithPoolSize(config.WORKER),
		parallelizer.WithJobQueueSize(config.QUEUE_SIZE),
	)
	defer group.Close()
	for p := 2; p <= pages; p++ {
		p := p // move to thread local memory
		group.Add(func() {
			crawlWithPage(queryURL, p)
		})
	}

	group.Wait()
}

func crawlWithPage(queryURL string, page int) {
	fullURL := fmt.Sprintf(queryURL+"&page=%d", page)
	var resp Resp
	c := util.CommonCollector()
	c.OnResponse(func(r *colly.Response) {
		json.Unmarshal(r.Body, &resp)
	})

	if config.DEBUG {
		fmt.Println("visiting", fullURL)
	}
	c.Visit(fullURL)

	model.CreateRepoBatch(resp.Repos)
}
