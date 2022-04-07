package main

import (
	"GitHunter/config"
	"GitHunter/model"
	"GitHunter/repo"
	"GitHunter/web"
    "os"
    "fmt"
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
    
	web.Crawl()
	repo.Clone()
}
