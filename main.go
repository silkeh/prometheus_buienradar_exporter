package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var regions, listenAddress, metricsPath string

	flag.StringVar(&listenAddress, "web.listen-address", ":9221", "Address to listen on for web interface and telemetry.")
	flag.StringVar(&metricsPath, "web.telemetry-path", "/metrics", "Path under which to expose metrics")
	flag.StringVar(&regions, "regions", "", "Regions to report separated by ','")
	flag.Parse()

	rs := strings.Split(regions, ",")
	if len(rs) == 0 || rs[0] == "" {
		log.Fatal("No regions configured.")
	}

	e := NewExporter(rs)

	prometheus.MustRegister(e)

	log.Println("Listening on", listenAddress)
	http.Handle(metricsPath, promhttp.Handler())
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
