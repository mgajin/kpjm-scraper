package crawler

import (
	"github.com/mgajin/kpjm-scraper/common/mq"
	"github.com/mgajin/kpjm-scraper/kpjm"
)

// Controller struct is used for starting crawlers.
type Controller struct {
	MessageHandler mq.MessageHandler
	ReceiveChan    chan interface{}
	SendChan       chan interface{}

	Done chan interface{}
}

// ControllerConfig struct holds controllers configuration.
type ControllerConfig struct {
	Service     *kpjm.CrawlingService
	ReceiveChan chan interface{}
	SendChan    chan interface{}
}

// NewController returns new controller.
func NewController(config *ControllerConfig, newMessageHandler Constructor) *Controller {

	crawlerConfig := &Config{
		Service:     config.Service,
		ReceiveChan: config.ReceiveChan,
		SendChan:    config.SendChan,
	}

	messageHandler := newMessageHandler(crawlerConfig)

	return &Controller{
		MessageHandler: messageHandler,
		ReceiveChan:    config.ReceiveChan,
		SendChan:       config.SendChan,
	}
}

// Crawl is start method for controller.
// Controller reads messages from receive channel and starts crawlers in separate go routines.
func (c *Controller) Crawl() {

	for message := range c.ReceiveChan {
		go c.MessageHandler.HandleMessage(message)
	}

	c.Done <- true
}
