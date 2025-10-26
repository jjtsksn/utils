package logger

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/thisismz/zerolog-loki/client"
)

var (
	once    sync.Once
	logger  *zerolog.Logger
	writers []io.Writer
	mu      sync.Mutex
)

func New(cfg Config) (*zerolog.Logger, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	var err error
	once.Do(func() {
		writers = make([]io.Writer, 0)
		if cfg.EnableConsole {
			consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: cfg.TimeFormat}
			writers = append(writers, consoleWriter)
		}

		if cfg.EnableLoki {
			labels := make(map[string]interface{})
			for _, label := range cfg.StaticLabels {
				labels[label] = "job"
			}
			var lokiWriter io.Writer
			lokiWriter, err = client.NewSimpleClient(cfg.LokiURL, cfg.LokiUser, cfg.LokiPass, client.WithStaticLabels(labels))
			if err != nil {
				return
			}
			writers = append(writers, lokiWriter)
		}
		multiWriter := zerolog.MultiLevelWriter(writers...)
		lg := zerolog.New(multiWriter).Level(cfg.LogLevel).With().Timestamp().Logger()
		logger = &lg

	})

	if err != nil {
		return nil, fmt.Errorf("create loki client: %w", err)
	}

	return logger, nil
}

func Close() {
	mu.Lock()
	defer mu.Unlock()
	for _, writer := range writers {
		if closer, ok := writer.(io.Closer); ok {
			_ = closer.Close()
		}
	}
}
