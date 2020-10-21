package handler

import (
	"fmt"

	"github.com/mgajin/kpjm-scraper/common/model"
	"github.com/mgajin/kpjm-scraper/common/mq"
	"github.com/mgajin/kpjm-scraper/storage/mongo"
)

type AdHandler struct {
	handler   *Handler
	adStorage *mongo.AdStorage
}

func NewAdHandler(config *Config, adStorage *mongo.AdStorage) *AdHandler {
	baseHandler := NewBaseHandler(config)

	return &AdHandler{
		handler:   baseHandler,
		adStorage: adStorage,
	}
}

func (ah *AdHandler) HandleMessage(message mq.Message) {
	ad := message.(*model.Ad)

	if err := ah.adStorage.StoreAd(ad); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Stored Ad[%d]\n", ad.ID)
}
