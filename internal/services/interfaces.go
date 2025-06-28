// Package services содержит интерфейсы для бизнес-логики приложения.
package services

import (
	fiber "github.com/gofiber/fiber/v2"
)

// UserServiceInterface определяет интерфейс для работы с пользователями.
type UserServiceInterface interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUserByUUID(ctx *fiber.Ctx) error
	UpdateUserByUUID(ctx *fiber.Ctx) error
	DeleteUserByUUID(ctx *fiber.Ctx) error
}
