package mq

// Message type represents message that will be passed through queues and channels.
type Message = interface{}

// MessageHandler handles a message consumed
// from the message queue.
type MessageHandler interface {
	HandleMessage(message Message)
}

// MessageBroker is used for consuming and publishing messages from and to message queue.
type MessageBroker interface {
	Consume() Message
	Publish(message Message)
}
