package main

import (
	"app/routes"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/user", routes.GetUser)
	v1.Post("/user", routes.CreateUser)
	v1.Put("/user", routes.UpdateUser)
	v1.Delete("/user", routes.DeleteUser)
}

func main() {
	_ = godotenv.Load(".env")
	app := fiber.New()
	app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.yml",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}))
	Routes(app)
	if err := app.Listen(os.Getenv("APP_PORT")); err != nil {
		return
	}
}
