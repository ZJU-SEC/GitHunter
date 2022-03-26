package main

import (
	"GitHunter/config"
	"GitHunter/model"
	"GitHunter/web"
)

func main() {
	config.Init()
	model.Init()
	web.Crawl()
}
