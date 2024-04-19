package main

import (
	"log-generator/worker"
)

func main() {
	var reporter chan string
	w := worker.NewWorker()
	w.LoadWorkers()
	<-reporter
}
