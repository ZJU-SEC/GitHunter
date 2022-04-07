package main

import (
	"GitHunter/config"
	"GitHunter/model"
	"GitHunter/repo"
	"GitHunter/web"
)

func main() {
	config.Init()
	model.Init()
	web.Crawl()
	repo.Clone()
}
