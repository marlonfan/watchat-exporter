package model

import "time"

// MysqlDriver mysql Driver config
type MysqlDriver struct {
	Source   string        `json:"source" yaml:"source" validate:"required"`
	Query    string        `mapstructure:"query" yaml:"query" validate:"required"`
	BindVar  []interface{} `mapstructure:"bind_var" yaml:"bind_var" validate:"required"`
	Interval string        `mapstructure:"interval" yaml:"interval" validate:"required"`
}

// PostgresDriver postgres Driver config
type PostgresDriver struct {
	Labels   map[string]string `mapstructure:"labels" yaml:"labels"`
	Help     string            `mapstructure:"help" yaml:"help"`
	Query    string            `mapstructure:"query" yaml:"query"`
	BindVar  string            `mapstructure:"bind_var" yaml:"bind_var"`
	Interval time.Duration     `mapstructure:"interval" yaml:"interval"`
}
