package web

import (
	"GitHunter/config"
	"GitHunter/model"
	"GitHunter/util"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"math"
)

type Resp struct {
	count int          `json:"total_count"`
	repos []model.Repo `json:"items"`
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

	c.Visit(fmt.Sprintf("https://api.github.com/search/repositories?q=language%3AJava+stars%3A%3E10000"))

	if resp.count == 0 {
		return
	} else if resp.count > 1000 {
		crawlWithOption((min+max)/2, max)
		crawlWithOption(min, (min+max)/2)
	}

	pages := resp.count/30 + 1

	// TODO put all items in resp into DB and iterate from page 2 to the final page

	for p := 2; p <= pages; p++ {

	}
}
