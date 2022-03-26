package web

import (
	"GitHunter/config"
	"GitHunter/model"
	"GitHunter/util"
	"encoding/json"
	"fmt"
	"math"

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
	crawlWithOption(MAX_STAR, math.MaxInt)
	crawlWithOption(config.MIN_STAR, MAX_STAR)
}

// For the cursor, this function return the repos having star in range [min, max]
func crawlWithOption(min, max int) {
	var resp Resp
	// send one request, if bigger than 1000, divide them into two parts
	c := util.CommonCollector()
	c.OnResponse(func(r *colly.Response) {
		json.Unmarshal(r.Body, &resp)
	})

	queryURL := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s+stars:%d..%d", config.LANGUAGE, min, max-1)
	fmt.Println(queryURL)
	c.Visit(queryURL)

	if resp.Count == 0 {
		return
	} else if resp.Count > 1000 {
		mid := (min + max) / 2
		crawlWithOption(mid, max)
		// solve duplication
		crawlWithOption(min, mid-1)
		return
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
