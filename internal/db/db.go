package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func Connect(connString string) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("pgx.ParseConfig: %v", err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, fmt.Errorf("pgx.NewWithConfig: %v", err)
	}

	return db, nil
}
