package elastic

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"log-generator/config"
	"log-generator/logger/generator"
	"net/http"
	"time"
)

type ElasticLogger struct {
	cfg config.LoggerConf
}

func (el ElasticLogger) Flush() {
	var docs []byte

	for j := 0; j < el.cfg.Documents; j++ {
		header := []byte(fmt.Sprintf(`{ "index" : { } }%s`, "\n"))
		docs = append(docs, header...)
		doc := generator.GenerateJSON()
		d, _ := json.Marshal(doc)
		docs = append(docs, d...)
		docs = append(docs, "\n"...)
	}

	// bod, err := json.Marshal(docs)

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s/_bulk", el.cfg.Host, el.cfg.Indexes[0]),
		bytes.NewReader(docs),
	)
	if err != nil {
		log.Printf("Error connecting: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(el.cfg.User, el.cfg.Password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := http.Client{
		Timeout:   el.cfg.OutTimeout * time.Second,
		Transport: tr,
	}

	before := time.Now()
	_, err = client.Do(req)
	if err != nil {
		log.Printf("Error posting: %s", err)
		return
	}

	after := time.Now().Sub(before)
	log.Printf("Posting took: %v", after)

	// Update the way latecy is being exported to prometheus metrics.
	// el.Exporter.Export("out", float64(after.Milliseconds()))
}

func NewElasticLogger(cfg config.LoggerConf) ElasticLogger {
	return ElasticLogger{cfg: cfg}
}
