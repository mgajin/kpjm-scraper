package crawler

import (
	"github.com/mgajin/kpjm-scraper/common/mq"
	"github.com/mgajin/kpjm-scraper/kpjm"
)

// Constructor is used for creating crawlers.
type Constructor func(config *Config) mq.MessageHandler

// Crawler struct
type Crawler struct {
	Service     *kpjm.CrawlingService
	ReceiveChan chan interface{}
	SendChan    chan interface{}
}

// Config struct is used for defining crawler's configuration.
type Config struct {
	Service     *kpjm.CrawlingService
	ReceiveChan chan interface{}
	SendChan    chan interface{}
}

// NewBaseCrawler returns instance of new Crawler.
// BaseCrawler is used in every crawler implementation.
func NewBaseCrawler(config *Config) *Crawler {
	return &Crawler{
		Service:     config.Service,
		ReceiveChan: config.ReceiveChan,
		SendChan:    config.SendChan,
	}
}
