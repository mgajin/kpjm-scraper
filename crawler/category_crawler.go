package crawler

import (
	"fmt"
	"log"

	"github.com/mgajin/kpjm-scraper/common/mq"
)

// CategoryCrawler struct.
// Crawler implementation for Categories.
type CategoryCrawler struct {
	*Crawler
}

// NewCategoryCrawler returns MessageHandler / CategoryCrawler.
func NewCategoryCrawler(config *Config) mq.MessageHandler {
	crawler := NewBaseCrawler(config)

	return &CategoryCrawler{
		Crawler: crawler,
	}
}

// HandleMessage is CategoryCrawler's implementation of MessageHandler interface.
func (cc *CategoryCrawler) HandleMessage(message mq.Message) {
	id := message.(int)
	page, pages, count := 1, 0, 0

	for {
		res, err := cc.Service.ScrapeCategory(id, page)
		if err != nil {
			log.Println(err)
			return
		}
		pages = res.Pages

		for _, ad := range res.Ads {
			cc.SendChan <- ad
			count++
		}

		fmt.Printf("Category[%d] - %d/%d - (%d)\n", id, page, pages, count)

		if page == pages || page > pages {
			break
		}

		page++
	}

	fmt.Printf("Finished category[%d] - ads(%d) - pages(%d)\n", id, count, pages)
}
