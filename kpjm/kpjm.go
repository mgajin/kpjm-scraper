package kpjm

import (
	"fmt"
	"strconv"

	"github.com/mgajin/kpjm-scraper/client"
	"github.com/mgajin/kpjm-scraper/common/model"
	"github.com/mgajin/kpjm-scraper/common/parser"
)

// CrawlingService is service that is used for crawling through categories and scraping ads.
type CrawlingService struct {
	Client     *client.KpjmClient
	ParserPool *parser.Pool
}

// NewCrawlingService returns instance of new CrawlingService.
func NewCrawlingService(kc *client.KpjmClient, pool *parser.Pool) *CrawlingService {
	return &CrawlingService{
		Client:     kc,
		ParserPool: pool,
	}
}

// ScrapeCategories fetches all categories and groups.
// Returns pointer slice of categories / groups and error.
func (service *CrawlingService) ScrapeCategories() (*[]model.Category, error) {
	req, err := client.NewRequest("action=categories")
	if err != nil {
		return nil, err
	}

	body, err := service.Client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	parserInstance := service.ParserPool.Get()

	categories, err := model.CategoriesFromJSON(body, parserInstance)
	if err != nil {
		return nil, err
	}

	service.ParserPool.Put(parserInstance)

	return categories, nil
}

// ScrapeCategory fetches category by ID and page number.
// Returns category filter and error.
func (service *CrawlingService) ScrapeCategory(categoryID int, page int) (*model.CategoryFilter, error) {
	route := fmt.Sprintf("v=1.0&data[category_id]=%s&action=search&data[page]=%s", strconv.Itoa(categoryID), strconv.Itoa(page))

	req, err := client.NewRequest(route)
	if err != nil {
		return nil, err
	}

	body, err := service.Client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	parserInstance := service.ParserPool.Get()

	category, err := model.CategoryFilterFromJSON(body, parserInstance)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// ScrapeAd fetches ad by ID.
// Returns ad and error.
func (service *CrawlingService) ScrapeAd(adID int) (*model.Ad, error) {
	route := fmt.Sprintf("ad_id=%s&v=2.0&action=ad", strconv.Itoa(adID))

	req, err := client.NewRequest(route)
	if err != nil {
		return nil, err
	}

	body, err := service.Client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	parserInstance := service.ParserPool.Get()

	ad, err := model.AdFromJSON(body, parserInstance)
	if err != nil {
		return nil, err
	}

	service.ParserPool.Put(parserInstance)

	return ad, nil
}
