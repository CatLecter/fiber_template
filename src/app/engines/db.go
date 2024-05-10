package engines

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"time"
)

type DBEngine struct {
	pool *pgxpool.Pool
}

func (db *DBEngine) createPool(ctx *context.Context) error {
	poolConfig, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"),
		),
	)
	if err != nil {
		log.Panicf("ERROR: unable to parse DATABASE_URL. %v", err)
	}
	poolConfig.MinConns = 0
	poolConfig.MaxConns = 20
	poolConfig.MaxConnIdleTime = 3 * time.Nanosecond
	poolConfig.MaxConnLifetime = 60 * time.Nanosecond
	db.pool, err = pgxpool.NewWithConfig(*ctx, poolConfig)
	if err != nil {
		log.Panicf("ERROR: unable to create connection pool. %v", err)
	}
	return nil
}

func (db *DBEngine) GetConn(ctx *context.Context) (*pgxpool.Conn, error) {
	if err := db.createPool(ctx); err != nil {
		log.Printf("ERROR: connection error. %v", err)
		return nil, err
	}
	conn, err := db.pool.Acquire(*ctx)
	if err != nil {
		// TODO: тут возникает ошибка в постгресе в PostgreSQL
		log.Printf("ERROR: acquiring a connection. %v", err)
		return nil, err
	}
	return conn, nil
}
