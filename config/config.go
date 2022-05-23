package config

// DriverConfig is driver config
type DriverConfig struct {
	Type       string            `json:"type" yaml:"type"`
	MetricType string            `json:"metric_type" yaml:"metric_type"`
	Name       string            `json:"name" yaml:"name"`
	Namespace  string            `json:"namespace" yaml:"namespace"` // metrics system is optional
	Subsystem  string            `json:"subsystem" yaml:"subsystem"` // metrics subsystem is optional
	Help       string            `json:"help" yaml:"help"`
	Labels     map[string]string `json:"labels" yaml:"labels"`
	Config     interface{}       `json:"config" yaml:"config"`
}

// SourceConfig is data source config
type SourceConfig struct {
	Name   string                 `json:"name" yaml:"name"`
	Type   string                 `json:"type" yaml:"type"`
	Config map[string]interface{} `json:"config" yaml:"config"`
}

// Config is the configuration for the application.
type Config struct {
	Lazy    bool
	Drivers []*DriverConfig
	Sources []*SourceConfig
}
