package handler

import "log"

type Handler struct {
	ReceiveChan chan interface{}
	SendChan    chan interface{}
	Logger      log.Logger
}

type Config struct {
	Logger log.Logger
}

func NewBaseHandler(config *Config) *Handler {
	return &Handler{}
}
