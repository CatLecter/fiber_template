package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) CreateUser(ctx *fiber.Ctx) error       { return h.service.CreateUser(ctx) }
func (h *Handler) GetUserByUUID(ctx *fiber.Ctx) error    { return h.service.GetUserByUUID(ctx) }
func (h *Handler) UpdateUserByUUID(ctx *fiber.Ctx) error { return h.service.UpdateUserByUUID(ctx) }
func (h *Handler) DeleteUserByUUID(ctx *fiber.Ctx) error { return h.service.DeleteUserByUUID(ctx) }
