package logger

import (
	"errors"
	"fmt"
	"os"

	"github.com/jjtsksn/utils/environment"
	"go.uber.org/zap"
)

type CancelFunc func() error

func New(env environment.Environment, opts ...zap.Option) (*zap.Logger, CancelFunc, error) {
	var cfg zap.Config
	options := []zap.Option{zap.AddCaller()}

	switch env {
	case environment.EnvLocal:
		cfg = newLocalConfig()
		options = append(options, zap.AddStacktrace(zap.WarnLevel))
	case environment.EnvDevelopment:
		cfg = newDevelopmentConfig()
		options = append(options, zap.AddStacktrace(zap.ErrorLevel))
	case environment.EnvProduction:
		cfg = newProductionConfig()
		options = append(options, zap.AddStacktrace(zap.FatalLevel))
	default:
		return nil, nil, fmt.Errorf("unknown environment: %s", env)
	}

	options = append(options, opts...)

	logger, err := newLogger(cfg, options...)
	if err != nil {
		return nil, nil, err
	}

	cancel := func() error { return safeSync(logger) }

	return logger, cancel, nil
}

func Must(appEnvironment environment.Environment, opts ...zap.Option) (*zap.Logger, CancelFunc) {
	logger, cancel, err := New(appEnvironment, opts...)
	if err != nil {
		panic(err)
	}
	return logger, cancel
}

func newLogger(cfg zap.Config, opts ...zap.Option) (*zap.Logger, error) {
	return cfg.Build(opts...)
}

func safeSync(logger *zap.Logger) error {
	if err := logger.Sync(); err != nil && !errors.Is(err, os.ErrInvalid) {
		return err
	}
	return nil
}
