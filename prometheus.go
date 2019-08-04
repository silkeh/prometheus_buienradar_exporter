package main

import (
	"log"
	"net/http"
	"time"

	"git.slxh.eu/prometheus/buienradar_exporter/buienradar"
	"github.com/prometheus/client_golang/prometheus"
)

// Namespace is the namespace for the Prometheus metrics
const Namespace = "buienradar"

// Metric represents an exported prometheus Metric
type Metric struct {
	*prometheus.GaugeVec
	GetValue func(s *buienradar.StationMeasurement) float64
}

// NewMetric return a new metric using the standard namespace and labels
func NewMetric(name, help string, getter func(s *buienradar.StationMeasurement) float64) *Metric {
	return &Metric{
		GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: Namespace,
			Name:      name,
			Help:      help,
		}, []string{"region"}),
		GetValue: getter,
	}
}

// Exporter is a Prometheus exporter
type Exporter struct {
	regions []string
	client  *buienradar.Client
	metrics []*Metric
}

// NewExporter initialises the Buienradar Prometheus exporter
func NewExporter(regions []string) *Exporter {
	return &Exporter{
		regions: regions,
		client:  buienradar.NewClient(&http.Client{Timeout: 1 * time.Second}),
		metrics: []*Metric{
			NewMetric(
				"temperature_celsius",
				"Temperature in °C",
				func(s *buienradar.StationMeasurement) float64 {
					return s.Temperature
				}),
			NewMetric(
				"feel_temperature_celsius",
				"Feel temperature in °C",
				func(s *buienradar.StationMeasurement) float64 {
					return s.Feeltemperature
				}),
			NewMetric(
				"pressure_hpa",
				"Atmospheric pressure in hPa",
				func(s *buienradar.StationMeasurement) float64 {
					return float64(s.Airpressure)
				}),
			NewMetric(
				"wind_mps",
				"Wind speed in m/s",
				func(s *buienradar.StationMeasurement) float64 {
					return s.Windspeed
				}),
		},
	}
}

// Describe describes all the Prometheus metrics
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range e.metrics {
		m.Describe(ch)
	}
}

// Collect collects data from the Buienradar API
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	data, err := e.client.Get()
	if err != nil {
		log.Printf("Error retrieving data: %s", err)
		return
	}

	for _, s := range data.Actual.Stationmeasurements {
		if contains(e.regions, s.Regio) {
			for _, m := range e.metrics {
				m.WithLabelValues(s.Regio).Set(m.GetValue(&s))
			}
		}
	}

	for _, m := range e.metrics {
		m.Collect(ch)
	}
}

// contains checks if a string slice contains a certain string
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
