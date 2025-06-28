// Package handlers реализует HTTP-обработчики для пользователей.
package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
)

// CreateUser создает нового пользователя.
func (h *Handler) CreateUser(ctx *fiber.Ctx) error { return h.service.CreateUser(ctx) }

// GetUserByUUID получает пользователя по UUID.
func (h *Handler) GetUserByUUID(ctx *fiber.Ctx) error { return h.service.GetUserByUUID(ctx) }

// UpdateUserByUUID обновляет пользователя по UUID.
func (h *Handler) UpdateUserByUUID(ctx *fiber.Ctx) error { return h.service.UpdateUserByUUID(ctx) }

// DeleteUserByUUID удаляет пользователя по UUID.
func (h *Handler) DeleteUserByUUID(ctx *fiber.Ctx) error { return h.service.DeleteUserByUUID(ctx) }
