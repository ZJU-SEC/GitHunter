package web

import (
	"GitHunter/config"
	"GitHunter/model"
	"GitHunter/util"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/shomali11/parallelizer"
	"math"
)

// Resp is a simple api response handler struct
type Resp struct {
	Count int          `json:"total_count"`
	Repos []model.Repo `json:"items"`
}

var MAX_STAR = 10000

func Crawl() {
	crawlWithOption(MAX_STAR, math.MaxInt)
	crawlWithOption(config.MIN_STAR, MAX_STAR)
}

func crawlWithOption(min, max int) {
	var resp Resp
	// send one request, if bigger than 1000, divide them into two parts
	c := util.CommonCollector()
	c.OnResponse(func(r *colly.Response) {
		json.Unmarshal(r.Body, &resp)
	})

	queryURL := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s+stars:%d", config.LANGUAGE)
	c.Visit(queryURL)

	if resp.Count == 0 {
		return
	} else if resp.Count > 1000 {
		crawlWithOption((min+max)/2, max)
		crawlWithOption(min, (min+max)/2)
	}

	model.CreateRepoBatch(resp.Repos)
	pages := (resp.Count-1)/30 + 1

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

func crawlWithPage(url string, page int) {
	// TODO request with page
}
