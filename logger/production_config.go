package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newProductionConfig() zap.Config {
	cfg := zap.NewProductionConfig()

	cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	cfg.Encoding = "json"
	cfg.Sampling = &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 100,
	}

	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime:    zapcore.EpochMillisTimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}

	return cfg
}
