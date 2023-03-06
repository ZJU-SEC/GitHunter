package main

import (
	"GitHunter/config"
	"GitHunter/model"
	"GitHunter/repo"
	"GitHunter/web"
	"fmt"
	"os"
)

func main() {
	config.Init()
	model.Init()

	if len(os.Args) < 2 {
		fmt.Println("require an argument")
	}

	switch os.Args[1] {
	case "clone":
		repo.Clone()
	case "crawl":
		web.Crawl()
	}
}
