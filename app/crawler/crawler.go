package crawler

import (
	"fmt"

	"github.com/mgajin/kpjm-scraper/client"
	"github.com/mgajin/kpjm-scraper/common/parser"
	crawlerService "github.com/mgajin/kpjm-scraper/crawler"
	"github.com/mgajin/kpjm-scraper/kpjm"
)

func InitializeChannels() (categories, ads, data chan interface{}) {
	categories = make(chan interface{})
	ads = make(chan interface{})
	data = make(chan interface{})

	return
}

func InitializeServices(client *client.KpjmClient) (services []*kpjm.CrawlingService) {

	for i := range defaultConfigQueue {
		parserPool := parser.NewPool()
		service := kpjm.NewCrawlingService(client, parserPool)
		services = append(services, service)
		fmt.Println("Initialized service ", i)
	}

	return
}

func InitializeControllers(services []*kpjm.CrawlingService) (controllers []*crawlerService.Controller) {

	for i, defaultConfig := range defaultConfigQueue {
		controller := initializeController(defaultConfig, services[i])
		controllers = append(controllers, controller)
		fmt.Println("Initialized controller ", i)
	}

	return
}

func initializeController(config *Config, service *kpjm.CrawlingService) (controller *crawlerService.Controller) {

	configuration := &crawlerService.ControllerConfig{
		Service:     service,
		ReceiveChan: config.ReceiveChan,
		SendChan:    config.SendChan,
	}

	controller = crawlerService.NewController(configuration, config.ServiceConstructor)
	return
}
