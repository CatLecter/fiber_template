// Package database содержит функции для работы с базой данных PostgreSQL.
package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

// MustNewPool создает новый пул соединений с базой данных PostgreSQL.
// Паникует при ошибке подключения.
func MustNewPool(
	postgresURI *string,
	maxConnections *int32,
	minConnections *int32,
	maxConnLifetime *time.Duration,
	maxConnIdleTime *time.Duration,
) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(*postgresURI)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err.Error())
	}
	config.MaxConns = *maxConnections
	config.MinConns = *minConnections
	config.MaxConnLifetime = *maxConnLifetime
	config.MaxConnIdleTime = *maxConnIdleTime
	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err.Error())
	}
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err.Error())
	}
	log.Info("Successfully connected to database")
	return pool
}
