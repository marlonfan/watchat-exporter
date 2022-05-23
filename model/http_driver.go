package model

import "time"

// HTTPDriver config
type HTTPDriver struct {
	Name     string            `mapstructure:"name" yaml:"name"`
	Labels   map[string]string `mapstructure:"labels" yaml:"labels"`
	Query    string            `mapstructure:"query" yaml:"query"`
	BindVar  string            `mapstructure:"bind_var" yaml:"bind_var"`
	Interval time.Duration     `mapstructure:"interval" yaml:"interval"`
}
