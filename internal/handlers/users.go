// Package handlers реализует HTTP-обработчики для пользователей.
package handlers

import (
	"fibertemplate/internal/services"
	fiber "github.com/gofiber/fiber/v2"
)

// UserHandler реализует интерфейс для работы с пользователями.
type UserHandler struct {
	service *services.Service
}

// CreateUser создает нового пользователя.
func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error { return h.service.CreateUser(ctx) }

// GetUserByUUID получает пользователя по UUID.
func (h *UserHandler) GetUserByUUID(ctx *fiber.Ctx) error { return h.service.GetUserByUUID(ctx) }

// UpdateUserByUUID обновляет пользователя по UUID.
func (h *UserHandler) UpdateUserByUUID(ctx *fiber.Ctx) error { return h.service.UpdateUserByUUID(ctx) }

// DeleteUserByUUID удаляет пользователя по UUID.
func (h *UserHandler) DeleteUserByUUID(ctx *fiber.Ctx) error { return h.service.DeleteUserByUUID(ctx) }
