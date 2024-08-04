package main

import (
	"github.com/CatLecter/gin_template/configs"
	"github.com/CatLecter/gin_template/internal/database"
	"github.com/CatLecter/gin_template/internal/handlers"
	"github.com/CatLecter/gin_template/internal/repositories"
	"github.com/CatLecter/gin_template/internal/services"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	cfg := configs.NewConfig()

	db, err := database.NewPool(
		&cfg.PostgresURI,
		&cfg.MaxConnections,
		&cfg.MinConnections,
		&cfg.MaxConnLifetime,
		&cfg.MaxConnIdleTime,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err.Error())
	}

	repos := repositories.NewRepository()
	service := services.NewService(cfg, db, repos)
	handler := handlers.NewHandler(service)

	app := handler.Routes()

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err.Error())
		return
	}
}
