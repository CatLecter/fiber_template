package routes

import (
	"app/service"
	"github.com/gofiber/fiber/v2"
)

func GetUser(ctx *fiber.Ctx) error { return service.GetUser(ctx) }

func CreateUser(ctx *fiber.Ctx) error { return service.CreateUser(ctx) }

func UpdateUser(ctx *fiber.Ctx) error { return service.UpdateUser(ctx) }

func DeleteUser(ctx *fiber.Ctx) error { return service.DeleteUser(ctx) }
