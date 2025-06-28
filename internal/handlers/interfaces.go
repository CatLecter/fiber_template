// Package handlers содержит интерфейсы для HTTP-обработчиков.
package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
)

// UserHandlerInterface определяет интерфейс для работы с пользователями.
type UserHandlerInterface interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUserByID(ctx *fiber.Ctx) error
	UpdateUserByID(ctx *fiber.Ctx) error
	DeleteUserByID(ctx *fiber.Ctx) error
}
