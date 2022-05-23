package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/marlonfan/watchat-exporter"
	"github.com/marlonfan/watchat-exporter/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v3"
)

func main() {
	var cfg *config.Config
	b, err := ioutil.ReadFile("../testdata/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		panic(err)
	}
	w, err := watchat.NewWatchat(cfg)
	if err != nil {
		panic(err)
	}
	w.Start()

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		http.ListenAndServe(":2112", nil)
	}()

	time.Sleep(time.Second * 30)
}
