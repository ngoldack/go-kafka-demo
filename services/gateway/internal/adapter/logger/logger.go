package logger

import (
	"fmt"
	"go-kafka-demo/services/gateway/internal/adapter/config"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func NewLogger(cfg config.LoggerConfig, env string) *slog.Logger {
	var level slog.Level
	var err error
	if level, err = parseLogLevel(cfg.Level); err != nil {
		slog.Default().Warn("failed to parse log level, using default level 'debug'", slog.Any("error", err))
		level = slog.LevelDebug
	}

	if env == "production" {
		return slog.New(slog.NewJSONHandler(
			os.Stderr,
			&slog.HandlerOptions{
				Level: level,
			},
		))
	}

	return slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level: level,
	}))
}

func parseLogLevel(level string) (slog.Level, error) {
	switch level {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	}

	return -1, fmt.Errorf("unknown log level: %s", level)
}
