package loki

import "log-generator/config"

type LokiLogger struct {
}

func NewLokiLogger(cfg config.LoggerConf) LokiLogger {
	return LokiLogger{}
}

func (ll LokiLogger) Flush() {

}
