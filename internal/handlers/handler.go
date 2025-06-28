// Package handlers реализует HTTP-обработчики для Fiber.
package handlers

import (
	"fibertemplate/internal/config"
	"fibertemplate/internal/services"
	"github.com/gofiber/contrib/swagger"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Handler содержит HTTP-обработчики приложения.
type Handler struct {
	service *services.Service
	UserHandlerInterface
}

// NewHandler создает новый экземпляр HTTP-обработчика.
func NewHandler(cfg *config.Config, db *pgxpool.Pool) *Handler {
	service := services.NewService(cfg, db)
	return &Handler{
		service:              service,
		UserHandlerInterface: &UserHandler{service: service},
	}
}

// Routes настраивает и возвращает маршруты приложения.
func (h *Handler) Routes() *fiber.App {
	app := fiber.New()

	app.Use(
		swagger.New(
			swagger.Config{
				BasePath: "/",
				FilePath: "./docs/api/swagger.yml",
				Path:     "docs",
				Title:    "Swagger API Docs",
			},
		),
	)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")
	user.Get("/", h.GetUserByUUID)
	user.Post("/", h.CreateUser)
	user.Put("/", h.UpdateUserByUUID)
	user.Delete("/", h.DeleteUserByUUID)

	return app
}
