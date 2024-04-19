package worker

import (
	"log"
	"log-generator/config"
	"log-generator/exporter"
	"log-generator/logger"
	"time"

	"github.com/go-co-op/gocron"
)

type Worker struct {
	cfg       config.Config
	Exporter  exporter.Exporter
	Logger    logger.Logger
	Scheduler *gocron.Scheduler
}

func NewWorker() Worker {
	cfg := config.NewConfig()
	l, err := logger.NewLogger(cfg.LoggerConfig)
	if err != nil {
		log.Println(err)
	}

	w := Worker{
		cfg:    cfg,
		Logger: l,
	}

	// w.Exporter = exporter.NewExporter(float64(w.InTimeout), float64(w.OutTimeout))
	return w
}

func (w Worker) launchWorker() {
	gc := gocron.NewScheduler(time.UTC)
	w.Scheduler = gc
	duration := w.cfg.WorkerConfig.Interval * time.Millisecond
	gc.Every(duration).Do(w.Logger.Flush)
	gc.StartAsync()
	gc.StartBlocking()
}

func (w Worker) query() {
	return
}

func (w Worker) LoadWorkers() {
	for i := 0; i < int(w.cfg.WorkerConfig.Workers); i++ {
		go w.launchWorker()
	}
}

func (w *Worker) Reset() {
	log.Println("Stopping")
	w.Scheduler.Stop()
	w.Scheduler.StopBlockingChan()
}
