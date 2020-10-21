package mq

type Message = interface{}

type Consumer interface {
	Consume() (<-chan Message, error)
}

type MessageHandler interface {
	HandleMessage(Message)
}
