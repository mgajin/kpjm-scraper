package handler

import (
	"github.com/mgajin/kpjm-scraper/common/mq"
	"github.com/mgajin/kpjm-scraper/controller/storage"
)

// Constructor is constructor for handlers.
// It is used for initializing handlers inside controller.
type Constructor func(config *Config) mq.MessageHandler

// Handler struct
type Handler struct {
	StoringService storage.StoringService
	PublishChan    chan interface{}
}

// Config struct holds Handler's configuration.
type Config struct {
	StoringService storage.StoringService
	PublishChan    chan interface{}
}

// NewBaseHandler returns new Handler.
// Base Handler is used in every handler implementation.
func NewBaseHandler(config *Config) *Handler {
	return &Handler{
		StoringService: config.StoringService,
		PublishChan:    config.PublishChan,
	}
}
