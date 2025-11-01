package logger

import (
	"log/slog"
	"os"
)

func New(env, serviceName string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case "prod":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
		}))
	case "dev":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}))
	case "local":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelDebug,
		}))
	default:
		logger = slog.Default()
	}
	return logger.With("service", serviceName)
}
