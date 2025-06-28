// Package handlers содержит интерфейсы для HTTP-обработчиков.
package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
)

// UserHandlerInterface определяет интерфейс для работы с пользователями.
type UserHandlerInterface interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUserByUUID(ctx *fiber.Ctx) error
	UpdateUserByUUID(ctx *fiber.Ctx) error
	DeleteUserByUUID(ctx *fiber.Ctx) error
}
