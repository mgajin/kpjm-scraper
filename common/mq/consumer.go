package mq

// Message type is used for message handlers.
type Message = interface{}

// MessageHandler interface
type MessageHandler interface {
	HandleMessage(Message)
}
