package config

import (
	"text/template"
)

// Logging configures logs options.
type Logging struct {
	Level string
	Color bool
}

// Format configures formatting options.
type Format struct {
	Date     string
	Template *template.Template
}

// Config reflects J's configuration.
type Config struct {
	Editor   string
	Home     string
	FileName string

	Log    *Logging
	Format *Format
}

// New returns a new configuration.
func New() *Config {
	return &Config{}
}
