package storage

import (
	"github.com/mgajin/kpjm-scraper/storage/handler"
	"github.com/mgajin/kpjm-scraper/storage/mongo"
)

const (
	databaseURI  = "mongodb+srv://mgajin:mgajin123@markogajin.zzsho.mongodb.net/kpjm?retryWrites=true&w=majority"
	databaseName = "kpjm"
)

var MongoConfig = &mongo.Config{
	ConnectionURL: databaseURI,
	DatabaseName:  databaseName,
}

var (
	AdHandlerConfig       = &handler.Config{}
	CategoryHandlerConfig = &handler.Config{}
)

var defaultConfigQueue []*handler.ControllerConfig

func SetDefaultConfigs(configs ...*handler.ControllerConfig) {
	defaultConfigQueue = append(defaultConfigQueue, configs...)
}
