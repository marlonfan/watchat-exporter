package metrics

import "github.com/prometheus/client_golang/prometheus"

// NewCounter returns a new Counter with the given name and help string.
func NewCounter(name, namespace, Subsystem string, labels map[string]string) prometheus.Counter {
	c := prometheus.NewCounter(prometheus.CounterOpts{
		Name:        name,
		Namespace:   namespace,
		Subsystem:   Subsystem,
		ConstLabels: labels,
	})
	_ = prometheus.Register(c)
	return c
}

// NewGauge returns a new Gauge with the given name and help string.
func NewGauge(name, namespace, Subsystem string, labels map[string]string) prometheus.Gauge {
	g := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        name,
		Namespace:   namespace,
		Subsystem:   Subsystem,
		ConstLabels: labels,
	})
	_ = prometheus.Register(g)
	return g
}
