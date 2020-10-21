package main

import (
	"github.com/mgajin/kpjm-scraper/app"
)

var done chan interface{}

func init() {
	done = make(chan interface{})
}

func main() {

	a := app.NewApp()

	go a.FetchCategories()

	for _, controller := range a.StorageControllers {
		go controller.Handle()
	}

	for _, controller := range a.CrawlerControllers {
		go controller.Crawl()
	}

	<-done
}
