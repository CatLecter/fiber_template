package handlers

import (
	"github.com/CatLecter/gin_template/internal/services"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

type Handler struct{ service *services.Service }

func NewHandler(service *services.Service) *Handler { return &Handler{service: service} }

func (h *Handler) Routes() *fiber.App {
	app := fiber.New()

	app.Use(
		swagger.New(
			swagger.Config{
				BasePath: "/",
				FilePath: "./api/swagger.yml",
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
