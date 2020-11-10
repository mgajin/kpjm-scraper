package async

import (
	"fmt"

	"github.com/mgajin/kpjm-scraper/common/model"
	"github.com/mgajin/kpjm-scraper/common/mq"
	"github.com/mgajin/kpjm-scraper/controller/handler"
	"github.com/pkg/errors"
)

// AdHandler struct is Handler implementation for ads.
type AdHandler struct {
	*handler.Handler
}

// NewAdHandler returns new AdHandler.
func NewAdHandler(config *handler.Config) mq.MessageHandler {

	baseHandler := handler.NewBaseHandler(config)

	return &AdHandler{baseHandler}
}

// HandleMessage is MessageHandler implementation for AdHandler.
// Receives Ad as message from controller and stores it in database.
func (ah *AdHandler) HandleMessage(message mq.Message) {

	ad, ok := message.(*model.Ad)

	if !ok {
		err := errors.New(fmt.Sprintf("ad handler:unknown message type %T", message)).Error()
		fmt.Println(err)
		return
	}

	if err := ah.StoringService.Store(ad); err != nil {
		fmt.Println(err) // handle error ?
		return
	}

	ah.PublishChan <- ad
}
