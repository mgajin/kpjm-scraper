package handler

import (
	"fmt"

	"github.com/mgajin/kpjm-scraper/common/model"
	"github.com/mgajin/kpjm-scraper/common/mq"
	"github.com/mgajin/kpjm-scraper/storage/mongo"
)

type CategoryHandler struct {
	handler         *Handler
	categoryStorage *mongo.CategoryStorage
}

func NewCategoryHandler(config *Config, adStorage *mongo.CategoryStorage) *CategoryHandler {
	baseHandler := NewBaseHandler(config)

	return &CategoryHandler{
		handler:         baseHandler,
		categoryStorage: adStorage,
	}
}

func (ch *CategoryHandler) HandleMessage(message mq.Message) {
	category := message.(*model.Category)

	if err := ch.categoryStorage.StoreCategory(category); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Stored Category[%d]\n", category.ID)
}
