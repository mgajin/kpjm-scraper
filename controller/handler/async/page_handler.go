package async

import (
	"fmt"

	"github.com/mgajin/kpjm-scraper/common/model"
	"github.com/mgajin/kpjm-scraper/common/mq"
	"github.com/mgajin/kpjm-scraper/controller/handler"
	"github.com/pkg/errors"
)

// PageHandler struct is Handler implementation for pages.
type PageHandler struct {
	*handler.Handler
}

// NewPageHandler returns new PageHandler.
func NewPageHandler(config *handler.Config) mq.MessageHandler {

	baseHandler := handler.NewBaseHandler(config)

	return &PageHandler{baseHandler}
}

// HandleMessage is MessageHandler implementation for PageHandler.
// Receives Page as message from controller and extracts ad ID's.
func (ph *PageHandler) HandleMessage(message mq.Message) {

	page, ok := message.(*model.Page)

	if !ok {
		err := errors.New(fmt.Sprintf("page handler:unknown message type %T", message)).Error()
		fmt.Println(err)
		return
	}

	for _, ad := range page.Ads {
		ph.PublishChan <- ad
	}
}
