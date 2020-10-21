package handler

import (
	"github.com/mgajin/kpjm-scraper/common/mq"
	"github.com/mgajin/kpjm-scraper/storage/mongo"
)

type Controller struct {
	MessageHandler mq.MessageHandler
	Storage        *mongo.Storage
	ReceiveChan    chan interface{}
	Done           chan interface{}
}

type ControllerConfig struct {
	Storage     *mongo.Storage
	ReceiveChan chan interface{}
}

func NewController(config *ControllerConfig, messageHandler mq.MessageHandler) *Controller {
	return &Controller{
		MessageHandler: messageHandler,
		Storage:        config.Storage,
		ReceiveChan:    config.ReceiveChan,
	}
}

func (c *Controller) Handle() {

	for message := range c.ReceiveChan {
		go c.MessageHandler.HandleMessage(message)
	}

	c.Done <- true
}
