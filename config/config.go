package config

import (
	"os"
	"strconv"
	"time"
)

type LoggerConf struct {
	Documents  int
	Host       string
	User       string
	Password   string
	Filepath   []string
	Indexes    []string
	OutTimeout time.Duration
	InTimeout  time.Duration
	LoggerType string
}

type WorkerConf struct {
	Workers  int
	Interval time.Duration
}

type Config struct {
	LoggerConfig LoggerConf
	WorkerConfig WorkerConf
}

func NewConfig() Config {
	// h := os.Getenv("HOST")
	// wrk := os.Getenv("WORKERS")
	// workers, _ := strconv.Atoi(wrk)
	// host := fmt.Sprintf("https://%v", h)
	return newDefaultConfig()
}

func newDefaultConfig() Config {
	return Config{
		LoggerConfig: LoggerConf{
			Documents:  1000,
			Host:       "localhost:9200",
			User:       "admin",
			Password:   "admin",
			Filepath:   []string{"/var/log/nginx/loki_testing_access.log"},
			Indexes:    []string{"kidred-test-index"},
			LoggerType: "file",
		},
		WorkerConfig: newWorkerConfig(),
	}
}

func newWorkerConfig() WorkerConf {
	wrk := os.Getenv("WORKERS")
	workers, _ := strconv.Atoi(wrk)

	intv := os.Getenv("INTERVAL")
	interval, _ := strconv.Atoi(intv)
	return WorkerConf{
		Workers:  workers,
		Interval: time.Duration(interval),
	}
}
