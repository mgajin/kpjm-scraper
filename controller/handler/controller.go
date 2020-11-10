package handler

import (
	"github.com/Jeffail/tunny"
	"github.com/mgajin/kpjm-scraper/common/mq"
	"github.com/mgajin/kpjm-scraper/controller/storage"
	"github.com/pkg/errors"
)

// Controller struct is used for controlling application flow between handlers.
type Controller struct {
	MessageHandler mq.MessageHandler
	MessageBroker  mq.MessageBroker
	Pool           *tunny.Pool
	ConsumeChan    chan interface{}
	PublishChan    chan interface{}
	Done           chan interface{}
}

// ControllerConfig struct holds configurations for controllers.
type ControllerConfig struct {
	StoringService  storage.StoringService
	MessageBroker   mq.MessageBroker
	ConsumeChan     chan interface{}
	PublishChan     chan interface{}
	PoolWorkerCount int
}

// PayloadWrapper struct wraps message handler and message.
// It is used for tunny pool.
type PayloadWrapper struct {
	MessageHandler mq.MessageHandler
	Message        mq.Message
}

// NewController creates new controller and initializes message handler for given handler constructor.
// Returns *Controller
func NewController(config *ControllerConfig, newMessageHandler Constructor) *Controller {

	handlerConfig := &Config{
		StoringService: config.StoringService,
		PublishChan:    config.PublishChan,
	}

	messageHandler := newMessageHandler(handlerConfig)

	pool := tunny.NewFunc(config.PoolWorkerCount, func(payload interface{}) interface{} {
		payloadWrapper, ok := payload.(PayloadWrapper)

		if !ok {
			return errors.New("couldn't convert payload")
		}

		payloadWrapper.MessageHandler.HandleMessage(payloadWrapper.Message)

		return nil
	})

	controller := &Controller{
		MessageHandler: messageHandler,
		MessageBroker:  config.MessageBroker,
		Pool:           pool,
		ConsumeChan:    config.ConsumeChan,
		PublishChan:    config.PublishChan,
		Done:           nil,
	}

	return controller
}

// Handle is start method for controller.
// Controller reads message from queue and starts handlers in separate go routines.
func (c *Controller) Handle() {

	go c.consumeMessages()
	go c.publishMessages()

	for message := range c.ConsumeChan {
		go c.Pool.Process(PayloadWrapper{
			MessageHandler: c.MessageHandler,
			Message:        message,
		})
	}

	c.Done <- true
}

// consumeMessages consumes messages through message broker.
// Sends consumed messages to channel.
func (c *Controller) consumeMessages() {

	for {
		message := c.MessageBroker.Consume()
		c.ConsumeChan <- message
	}
}

// publishMessages reads messages from channel.
// Message broker publishes messages.
func (c *Controller) publishMessages() {

	for message := range c.PublishChan {
		c.MessageBroker.Publish(message)
	}
}
