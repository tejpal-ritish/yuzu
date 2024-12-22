package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost/chatapp?sslmode=disable")
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
