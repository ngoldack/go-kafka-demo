package main

import (
	"context"
	"go-kafka-demo/services/gateway/internal/adapter/config"
	"go-kafka-demo/services/gateway/internal/adapter/logger"
	"go-kafka-demo/services/gateway/internal/adapter/persistance/mariadb"
	"log/slog"
	"os"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	logger := logger.NewLogger(cfg.Logger, cfg.App.Env)
	slog.SetDefault(logger)

	logger.InfoContext(ctx, "Starting the application",
		slog.String("app", cfg.App.Name),
		slog.String("env", cfg.App.Env),
		slog.String("version", cfg.App.Version),
	)

	db, err := mariadb.NewDB(ctx, &cfg.MariaDb)
	if err != nil {
		slog.ErrorContext(ctx, "Error initilising db connection", slog.Any("error", err))
		return err
	}
	defer db.Close()
	slog.InfoContext(ctx, "Successfully connected to db")

	err = db.Migrate(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Error migrating db", slog.Any("error", err))
		return err
	}
	slog.InfoContext(ctx, "Successfully migrated db")

	// Dependency Injection
	// Job
	

	return nil
}
