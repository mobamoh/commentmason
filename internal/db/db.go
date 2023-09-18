package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)
	dbConn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %w", err)
	}
	return &Database{
		Client: dbConn,
	}, nil
}

func (db *Database) Ping(ctx context.Context) error {
	return db.Client.DB.PingContext(ctx)
}
