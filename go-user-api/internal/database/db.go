package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres://postgres:postgres@localhost:5432/users?sslmode=disable"
	}

	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect database: %w", err)
	}
	return pool, nil
}