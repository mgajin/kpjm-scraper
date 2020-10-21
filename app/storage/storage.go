package storage

import (
	"fmt"
	"log"

	"github.com/mgajin/kpjm-scraper/common/mq"
	"github.com/mgajin/kpjm-scraper/storage/handler"
	"github.com/mgajin/kpjm-scraper/storage/mongo"
)

func InitializeStorage(config *mongo.Config) *mongo.Storage {

	storage, err := mongo.NewStorage(config)
	if err != nil {
		log.Fatal(err)
	}

	return storage
}

func InitializeControllers(messageHandlers ...mq.MessageHandler) (controllers []*handler.Controller) {

	for i, config := range defaultConfigQueue {
		controller := initializeController(config, messageHandlers[i])
		controllers = append(controllers, controller)
		fmt.Printf("Storage controller[%d] initialized.\n", i)
	}

	return
}

func initializeController(config *handler.ControllerConfig, messageHandler mq.MessageHandler) (controller *handler.Controller) {

	// configuration := &handler.ControllerConfig{
	// 	Storage:     config.Storage,
	// 	ReceiveChan: config.ReceiveChan,
	// }

	controller = handler.NewController(config, messageHandler)
	return
}
