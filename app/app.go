package app

import (
	"log"

	"github.com/mgajin/kpjm-scraper/app/crawler"
	"github.com/mgajin/kpjm-scraper/app/storage"
	"github.com/mgajin/kpjm-scraper/client"
	crawlerService "github.com/mgajin/kpjm-scraper/crawler"
	"github.com/mgajin/kpjm-scraper/kpjm"
	"github.com/mgajin/kpjm-scraper/storage/handler"
	"github.com/mgajin/kpjm-scraper/storage/mongo"
)

var (
	categories, ads, data chan interface{}
)

type App struct {
	CrawlerService     *kpjm.CrawlingService
	CrawlerControllers []*crawlerService.Controller
	StorageControllers []*handler.Controller
}

func NewApp() *App {

	// Message channels
	categories, ads, data = crawler.InitializeChannels()

	// Crawler Controller configurations
	categoryConfig := crawler.NewConfig(crawlerService.NewCategoryCrawler, categories, ads)
	adConfig := crawler.NewConfig(crawlerService.NewAdCrawler, ads, data)

	// Set crawler default configurations
	crawler.DefaultConfigQueue(categoryConfig, adConfig)

	// Initialize Crawling services and controllers
	kpjmClient := client.New()
	services := crawler.InitializeServices(kpjmClient)
	crawlerControllers := crawler.InitializeControllers(services)

	// Storage configurations
	storageConfig := storage.MongoConfig
	mongoStorage := storage.InitializeStorage(storageConfig)
	adStorage := mongo.NewAdStorage(mongoStorage)
	categoryStorage := mongo.NewCategoryStorage(mongoStorage)

	// Storage Handler and Controller configurations
	adHandler := handler.NewAdHandler(storage.AdHandlerConfig, adStorage)
	categoryHandler := handler.NewCategoryHandler(storage.CategoryHandlerConfig, categoryStorage)
	adStorageConfig := &handler.ControllerConfig{
		Storage:     mongoStorage,
		ReceiveChan: data,
	}
	categoryStorageConfig := &handler.ControllerConfig{
		Storage:     mongoStorage,
		ReceiveChan: nil,
	}

	// Initialize storage controllers
	storage.SetDefaultConfigs(adStorageConfig, categoryStorageConfig)
	storageControllers := storage.InitializeControllers(adHandler, categoryHandler)

	return &App{
		CrawlerService:     services[0],
		StorageControllers: storageControllers,
		CrawlerControllers: crawlerControllers,
	}
}

func (a *App) FetchCategories() {
	response, err := a.CrawlerService.ScrapeCategories()

	if err != nil {
		log.Fatal(err)
	}

	for _, category := range *response {
		// if i == 2 {
		// 	break
		// }
		categories <- category.ID
	}
}
