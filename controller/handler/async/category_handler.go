package async

import (
	"fmt"

	"github.com/mgajin/kpjm-scraper/common/model"
	"github.com/mgajin/kpjm-scraper/common/mq"
	"github.com/mgajin/kpjm-scraper/controller/handler"
	"github.com/pkg/errors"
)

// CategoryHandler struct is Handler implementation for categories.
type CategoryHandler struct {
	*handler.Handler
}

// NewCategoryHandler returns new CategoryHandler.
func NewCategoryHandler(config *handler.Config) mq.MessageHandler {

	baseHandler := handler.NewBaseHandler(config)

	return &CategoryHandler{baseHandler}
}

// HandleMessage is MessageHandler implementation for CategoryHandler.
// Receives Category as message from controller and stores it in database.
// Publishes tasks for scraping first pages of each category.
func (ch *CategoryHandler) HandleMessage(message mq.Message) {

	category, ok := message.(*model.Category)

	if !ok {
		err := errors.New(fmt.Sprintf("category handler:unknown message type %T", message)).Error()
		fmt.Println(err)
		return
	}

	// Commented for testing
	// if err := ch.StoringService.Store(category); err != nil {
	// 	fmt.Println(err) // handle error ?
	// 	return
	// }

	// if category parent is not 0 than it is group and needs to be published.
	if category.Parent > 0 {
		task := model.NewPageQuery(1, category.Parent, category.ID)
		ch.PublishChan <- task
	}
}
