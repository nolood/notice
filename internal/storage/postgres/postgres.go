package postgres

import (
	"context"
	"fmt"
	"notification-service/internal/config"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg *config.Config) *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), cfg.Storage.DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
