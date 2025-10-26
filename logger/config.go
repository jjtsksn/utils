package logger

import (
	"errors"

	"github.com/rs/zerolog"
)

type Config struct {
	StaticLabels  map[string]any
	LokiURL       string
	LokiUser      string
	LokiPass      string
	LogLevel      zerolog.Level
	TimeFormat    string
	EnableLoki    bool
	EnableConsole bool
}

func (c *Config) Validate() error {
	if c.StaticLabels == nil || len(c.StaticLabels) == 0 {
		return errors.New("StaticLabels must not be empty")
	}
	if c.TimeFormat == "" {
		return errors.New("time format is required")
	}
	if c.EnableLoki && c.LokiURL == "" {
		return errors.New("loki url is required when loki is enabled")
	}
	return nil
}
