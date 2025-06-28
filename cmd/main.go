// Package main запускает HTTP сервер Fiber.
package main

import (
	cfg "fibertemplate/internal/config"
	"fibertemplate/internal/database"
	"fibertemplate/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	cfg := cfg.NewConfig()

	db := database.MustNewPool(
		&cfg.PostgresURI,
		&cfg.MaxConnections,
		&cfg.MinConnections,
		&cfg.MaxConnLifetime,
		&cfg.MaxConnIdleTime,
	)

	handler := handlers.NewHandler(cfg, db)

	app := handler.Routes()

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err.Error())
		return
	}
}
