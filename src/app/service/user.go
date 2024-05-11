package service

import (
	"app/dto"
	"app/repositories"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
)

func GetUser(ctx *fiber.Ctx) error {
	userID, err := uuid.Parse(ctx.Query("user_id"))
	if err != nil {
		log.Printf("ERROR: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "user ID parsing error"})
	}
	user, err := repositories.GetUserByID(&userID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(
			fiber.Map{"result": fmt.Sprintf("user with ID=%v not found", userID)},
		)
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func CreateUser(ctx *fiber.Ctx) error {
	requestUser := new(dto.RequestUserDTO)
	if err := ctx.BodyParser(&requestUser); err != nil {
		log.Printf("ERROR: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "cannot parse JSON"})
	}
	responseUser, err := repositories.CreateUser(requestUser)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "error when creating a user"})
	}
	return ctx.Status(fiber.StatusOK).JSON(responseUser)
}

func UpdateUser(ctx *fiber.Ctx) error {
	userID, err := uuid.Parse(ctx.Query("user_id"))
	if err != nil {
		log.Printf("ERROR: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "user ID parsing error"})
	}
	_, err = repositories.GetUserByID(&userID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(
			fiber.Map{"result": fmt.Sprintf("user with ID=%v not found", userID)},
		)
	}
	requestUser := new(dto.RequestUserDTO)
	if err := ctx.BodyParser(&requestUser); err != nil {
		log.Printf("ERROR: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "cannot parse JSON"})
	}
	responseUser, err := repositories.UpdateUserByID(&userID, requestUser)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "error when update a user"})
	}
	return ctx.Status(fiber.StatusOK).JSON(responseUser)
}

func DeleteUser(ctx *fiber.Ctx) error {
	var userID uuid.UUID = uuid.MustParse(ctx.Query("user_id"))
	err := repositories.DeleteUserByID(&userID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(
			fiber.Map{"result": fmt.Sprintf("user with ID=%v not found", userID)},
		)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"result": "OK"})
}
