package async

import (
	"fmt"

	"github.com/mgajin/kpjm-scraper/common/model"
	"github.com/mgajin/kpjm-scraper/common/mq"
	"github.com/mgajin/kpjm-scraper/controller/handler"
	"github.com/pkg/errors"
)

// TaskHandler struct is handler for generating page queries.
type TaskHandler struct {
	*handler.Handler
}

// NewTaskHandler returns new TaskHandler.
func NewTaskHandler(config *handler.Config) mq.MessageHandler {

	baseHandler := handler.NewBaseHandler(config)

	return &TaskHandler{baseHandler}
}

// HandleMessage is MessageHandler implementation for TaskHandler.
// Receives first Page as message and generates PageQuery task for each page.
func (th *TaskHandler) HandleMessage(message mq.Message) {

	page, ok := message.(*model.Page)

	if !ok {
		err := errors.New(fmt.Sprintf("task handler:unknown message type %T", message)).Error()
		fmt.Println(err)
		return
	}

	for i := 1; i <= page.Pages; i++ {
		task := model.NewPageQuery(i, page.Category, page.Group)
		th.PublishChan <- task
	}
}
