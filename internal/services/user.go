package services

import (
	"fmt"
	"github.com/CatLecter/gin_template/internal/repositories"
	"github.com/CatLecter/gin_template/internal/schemes"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	repo repositories.UserRepositoryInterface
}

func NewUserService(repo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{repo: repo}
}

func (srv *UserService) GetUserByUUID(ctx *fiber.Ctx) error {
	userUUID, err := uuid.Parse(ctx.Query("uuid"))
	if err != nil {
		log.Errorf("Error parsing user UUID: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "cannot parse user UUID"})
	}
	userResp, err := srv.repo.GetUserByUUID(&userUUID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(
			fiber.Map{"result": fmt.Sprintf("user with UUID=%v not found", userUUID)},
		)
	}
	return ctx.Status(fiber.StatusOK).JSON(userResp)
}

func (srv *UserService) CreateUser(ctx *fiber.Ctx) error {
	user := schemes.UserRequest{}
	if err := ctx.BodyParser(&user); err != nil {
		log.Warnf("Error parsing body: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "cannot parse JSON"})
	}
	isExists, err := srv.repo.CheckUserByPhone(&user.Phone)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "error checking the user's existence"})
	}
	if *isExists {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			fiber.Map{"result": fmt.Sprintf("user with phone number %v already exists", user.Phone)},
		)
	}
	userResp, err := srv.repo.CreateUser(&user)
	if err != nil {
		log.Warnf("Error when creating a user: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "user cannot be created"})
	}
	return ctx.Status(fiber.StatusOK).JSON(userResp)
}

func (srv *UserService) UpdateUserByUUID(ctx *fiber.Ctx) error {
	userUUID, err := uuid.Parse(ctx.Query("uuid"))
	if err != nil {
		log.Warnf("Error parsing user UUID: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "cannot parse user UUID"})
	}
	user := schemes.UserRequest{}
	if err := ctx.BodyParser(&user); err != nil {
		log.Warnf("Error parsing body: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "cannot parse JSON"})
	}

	isExists, err := srv.repo.CheckUserByPhone(&user.Phone)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "error checking the user's existence"})
	}
	if *isExists {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			fiber.Map{"result": fmt.Sprintf("user with phone number %v already exists", user.Phone)},
		)
	}
	responseUser, err := srv.repo.UpdateUserByUUID(&userUUID, &user)
	if err != nil {
		log.Warnf("Error when updating a user: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "user cannot be updated"})
	}
	return ctx.Status(fiber.StatusOK).JSON(responseUser)
}

func (srv *UserService) DeleteUserByUUID(ctx *fiber.Ctx) error {
	userUUID, err := uuid.Parse(ctx.Query("uuid"))
	if err != nil {
		log.Warnf("Error parsing user UUID: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"result": "cannot parse user UUID"})
	}
	err = srv.repo.DeleteUserByUUID(&userUUID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(
			fiber.Map{"result": fmt.Sprintf("user with an UUID=%v cannot be deleted", userUUID)},
		)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"result": "success"})
}
