package crawler

import (
	crawlerService "github.com/mgajin/kpjm-scraper/crawler"
)

var defaultConfigQueue []*Config

type Config struct {
	QueueName          string
	ServiceConstructor crawlerService.Constructor
	ReceiveChan        chan interface{}
	SendChan           chan interface{}
}

func NewConfig(constructor crawlerService.Constructor, receiveChan, sendChan chan interface{}) *Config {
	return &Config{
		QueueName:          "",
		ServiceConstructor: constructor,
		ReceiveChan:        receiveChan,
		SendChan:           sendChan,
	}
}

func DefaultConfigQueue(configs ...*Config) {
	defaultConfigQueue = append(defaultConfigQueue, configs...)
}
