package crawler

import (
	"fmt"

	"github.com/mgajin/kpjm-scraper/common/mq"
)

// AdCrawler struct.
// Crawler implementation for Ads.
type AdCrawler struct {
	*Crawler
}

// NewAdCrawler returns AdCrawler / MessageHandler.
func NewAdCrawler(config *Config) mq.MessageHandler {
	crawler := NewBaseCrawler(config)

	return &AdCrawler{
		Crawler: crawler,
	}
}

// HandleMessage is AdCrawler's implementation of MessageHandler interface.
// It is used for scraping ad from website and sends it through channel to storage.
func (ac *AdCrawler) HandleMessage(message mq.Message) {
	id := message.(*int)

	ad, err := ac.Service.ScrapeAd(*id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if ad != nil {
		fmt.Printf("Scraped Ad[%v]\n", ad.ID)
		// ac.SendChan <- ad
	}
}
