package db

import (
	"context"
	"database/sql"
	"time"
)

func New(
	addr string,
	maxOpenConns int,
	maxIdleConns int,
	maxIdleTime string,
) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)

	if err != nil {
		return nil, err
	}

	// Configuration
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Ping SQL connection to verify if database is still alive
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
