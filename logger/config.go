package logger

import (
	"errors"

	"github.com/rs/zerolog"
)

type Config struct {
	StaticLabel   []string
	LokiURL       string
	LokiUser      string
	LokiPass      string
	LogLevel      zerolog.Level
	TimeFormat    string
	EnableLoki    bool
	EnableConsole bool
}

func (c *Config) Validate() error {
	if c.TimeFormat == "" {
		return errors.New("time format is required")
	}
	if c.EnableLoki && c.LokiURL == "" {
		return errors.New("loki url is required when loki is enabled")
	}
	return nil
}
