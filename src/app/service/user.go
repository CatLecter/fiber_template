package service

import (
	"app/repositories"
	"app/schemes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
)

func GetUser(ctx *fiber.Ctx) error {
	userUUID, err := uuid.Parse(ctx.Query("uuid"))
	if err != nil {
		log.Printf("ERROR: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "UUID parsing error"})
	}
	user, err := repositories.GetUserByUUID(&userUUID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(
			fiber.Map{"result": fmt.Sprintf("user with UUID=%v not found", userUUID)},
		)
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(schemes.UserReq)
	if err := ctx.BodyParser(&user); err != nil {
		log.Printf("ERROR: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "cannot parse JSON"})
	}
	userResp, err := repositories.CreateUser(user)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "error when creating a user"})
	}
	return ctx.Status(fiber.StatusOK).JSON(userResp)
}

func UpdateUser(ctx *fiber.Ctx) error {
	userUUID, err := uuid.Parse(ctx.Query("uuid"))
	if err != nil {
		log.Printf("ERROR: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "UUID parsing error"})
	}
	_, err = repositories.GetUserByUUID(&userUUID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(
			fiber.Map{"result": fmt.Sprintf("user with UUID=%v not found", userUUID)},
		)
	}
	user := new(schemes.UserReq)
	if err := ctx.BodyParser(&user); err != nil {
		log.Printf("ERROR: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "cannot parse JSON"})
	}
	userResp, err := repositories.UpdateUserByUUID(&userUUID, user)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "error when update a user"})
	}
	return ctx.Status(fiber.StatusOK).JSON(userResp)
}

func DeleteUser(ctx *fiber.Ctx) error {
	var userUUID uuid.UUID = uuid.MustParse(ctx.Query("uuid"))
	err := repositories.DeleteUserByUUID(&userUUID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(
			fiber.Map{"result": fmt.Sprintf("user with UUID=%v not found", userUUID)},
		)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"result": "OK"})
}
