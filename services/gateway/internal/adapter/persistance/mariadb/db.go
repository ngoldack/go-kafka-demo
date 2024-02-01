package mariadb

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"go-kafka-demo/services/gateway/internal/adapter/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type DB struct {
	db *sql.DB
}

func NewDB(ctx context.Context, cfg *config.MariaDbConfig) (*DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database))
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return &DB{
		db,
	}, nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) Migrate(ctx context.Context) error {
	goose.SetBaseFS(migrationsFS)
	if err := goose.SetDialect("mysql"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}
	if err := goose.Up(d.db, "migrations"); err != nil {
		return fmt.Errorf("failed to migrate db: %w", err)
	}

	return nil
}
