// Package handlers реализует HTTP-обработчики для пользователей.
package handlers

import (
	handlerutilities "fibertemplate/internal/handlers/utils"
	"fibertemplate/internal/schemes"
	"fibertemplate/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UserHandler реализует интерфейс для работы с пользователями.
type UserHandler struct {
	service *services.Service
}

// CreateUser создает нового пользователя.
func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var req schemes.UserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlerutilities.NewHTTPError("cannot parse JSON"))
	}
	user, err := h.service.CreateUser(ctx.Context(), &req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlerutilities.NewHTTPError(err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

// GetUserByID получает пользователя по ID.
func (h *UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	idStr := ctx.Query("id")
	userID, err := uuid.Parse(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlerutilities.NewHTTPError("cannot parse user ID"))
	}
	user, err := h.service.GetUserByID(ctx.Context(), userID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(handlerutilities.NewHTTPError(err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

// UpdateUserByID обновляет пользователя по ID.
func (h *UserHandler) UpdateUserByID(ctx *fiber.Ctx) error {
	idStr := ctx.Query("id")
	userID, err := uuid.Parse(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlerutilities.NewHTTPError("cannot parse user ID"))
	}
	var req schemes.UserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlerutilities.NewHTTPError("cannot parse JSON"))
	}
	user, err := h.service.UpdateUserByID(ctx.Context(), userID, &req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlerutilities.NewHTTPError(err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

// DeleteUserByID удаляет пользователя по ID.
func (h *UserHandler) DeleteUserByID(ctx *fiber.Ctx) error {
	idStr := ctx.Query("id")
	userID, err := uuid.Parse(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlerutilities.NewHTTPError("cannot parse user ID"))
	}
	err = h.service.DeleteUserByID(ctx.Context(), userID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(handlerutilities.NewHTTPError(err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(schemes.HTTPResponse{Result: "success", Msg: "user deleted successfully"})
}
