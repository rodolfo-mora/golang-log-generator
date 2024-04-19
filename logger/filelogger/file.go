package filelogger

import (
	"log"
	"log-generator/config"
	"log-generator/logger/generator"
	"os"
)

type FileLogger struct {
	cfg config.LoggerConf
}

func NewFileLogger(cfg config.LoggerConf) FileLogger {
	return FileLogger{
		cfg: cfg,
	}
}

func (fl FileLogger) Flush() {
	line := generator.GenerateString()
	logfile, err := openFile(fl.cfg.Filepath[0])
	if err != nil {
		log.Println(err)
	}
	defer logfile.Close()

	log.SetOutput(logfile)
	log.Println(line)
}

// Deprecated but useful
func writeToFile(cont chan string, f *os.File) {
	logger := log.New(f, "", log.LstdFlags)
	for line := range cont {
		logger.Output(2, line)
	}
}

func openFile(path string) (*os.File, error) {
	var file *os.File
	var err error

	file, err = os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return file, err
	}

	return file, err
}
