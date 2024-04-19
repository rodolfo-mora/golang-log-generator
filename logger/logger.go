package logger

import (
	"errors"
	"log"
	"log-generator/config"
	"log-generator/logger/elastic"
	"log-generator/logger/filelogger"
	"log-generator/logger/loki"
)

type Logger interface {
	Flush()
}

func NewLogger(cfg config.LoggerConf) (Logger, error) {
	switch cfg.LoggerType {
	case "loki":
		return loki.NewLokiLogger(cfg), nil
	case "elastic":
		return elastic.NewElasticLogger(cfg), nil
	case "file":
		log.Println("Loading file logger")
		return filelogger.NewFileLogger(cfg), nil
	}
	return nil, errors.New("Unsupported logger type selected")
}
