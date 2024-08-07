package services

import (
	"context"
	"fmt"
	"github.com/CatLecter/gin_template/configs"
	"github.com/CatLecter/gin_template/internal/repositories"
	"github.com/CatLecter/gin_template/internal/schemes"
	"github.com/CatLecter/gin_template/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	cfg  *configs.Config
	db   *pgxpool.Pool
	repo repositories.UserRepositoryInterface
}

func NewUserService(cfg *configs.Config, db *pgxpool.Pool, repo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{cfg, db, repo}
}

func (srv *UserService) GetUserByUUID(ctx *fiber.Ctx) error {
	var err error
	var userUUID uuid.UUID
	if userUUID, err = uuid.Parse(ctx.Query("uuid")); err != nil {
		log.Warnf("Error parsing user UUID: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewHTTPError("cannot parse user UUID"))
	}
	timeoutCtx, cancel := context.WithTimeout(ctx.Context(), srv.cfg.MaxConnLifetime)
	defer cancel()
	conn, err := srv.db.Acquire(timeoutCtx)
	defer conn.Release()
	if err != nil {
		log.Warnf("A new database connection could not be established: %v", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.NewHTTPError("a new database connection could not be established"),
		)
	}
	if userResp, err := srv.repo.GetUserByUUID(&timeoutCtx, conn, &userUUID); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(
			utils.NewHTTPError(fmt.Sprintf("user with UUID=%v not found", userUUID)),
		)
	} else {
		return ctx.Status(fiber.StatusOK).JSON(userResp)
	}
}

func (srv *UserService) CreateUser(ctx *fiber.Ctx) error {
	newUser := new(schemes.UserRequest)
	if err := ctx.BodyParser(&newUser); err != nil {
		log.Warnf("Error parsing body: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewHTTPError("cannot parse JSON"))
	}
	timeoutCtx, cancel := context.WithTimeout(ctx.Context(), srv.cfg.MaxConnLifetime)
	defer cancel()
	conn, err := srv.db.Acquire(timeoutCtx)
	defer conn.Release()
	if err != nil {
		log.Warnf("A new database connection could not be established: %v", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.NewHTTPError("a new database connection could not be established"),
		)
	}
	if isExists, err := srv.repo.CheckUserByPhone(&timeoutCtx, conn, &newUser.Phone); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.NewHTTPError("error checking the user's existence"),
		)
	} else if *isExists {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			utils.NewHTTPError(fmt.Sprintf("user with phone number %v already exists", newUser.Phone)),
		)
	}
	if userResp, err := srv.repo.CreateUser(&timeoutCtx, conn, newUser); err != nil {
		log.Warnf("Error when creating a user: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewHTTPError("user cannot be created"))
	} else {
		return ctx.Status(fiber.StatusOK).JSON(userResp)
	}
}

func (srv *UserService) UpdateUserByUUID(ctx *fiber.Ctx) error {
	var err error
	var userUUID uuid.UUID
	if userUUID, err = uuid.Parse(ctx.Query("uuid")); err != nil {
		log.Warnf("Error parsing user UUID: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewHTTPError("cannot parse user UUID"))
	}

	timeoutCtx, cancel := context.WithTimeout(ctx.Context(), srv.cfg.MaxConnLifetime)
	defer cancel()

	conn, err := srv.db.Acquire(timeoutCtx)
	defer conn.Release()
	if err != nil {
		log.Warnf("A new database connection could not be established: %v", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.NewHTTPError("a new database connection could not be established"),
		)
	}

	if isExistsUser, err := srv.repo.CheckUserByUUID(&timeoutCtx, conn, &userUUID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewHTTPError("error checking the user's existence"))
	} else if !*isExistsUser {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			utils.NewHTTPError(fmt.Sprintf("user with uuid %v not found", userUUID)),
		)
	}
	newUserData := new(schemes.UserRequest)
	if err := ctx.BodyParser(&newUserData); err != nil {
		log.Warnf("Error parsing body: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewHTTPError("cannot parse JSON"))
	}
	if isExistsPhone, err := srv.repo.CheckUserByPhone(&timeoutCtx, conn, &newUserData.Phone); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewHTTPError("error checking the user's existence"))
	} else if *isExistsPhone {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			utils.NewHTTPError(fmt.Sprintf("user with phone number %v already exists", newUserData.Phone)),
		)
	}
	if userResp, err := srv.repo.UpdateUserByUUID(&timeoutCtx, conn, &userUUID, newUserData); err != nil {
		log.Warnf("Error when updating a user: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewHTTPError("user cannot be updated"))
	} else {
		return ctx.Status(fiber.StatusOK).JSON(userResp)
	}
}

func (srv *UserService) DeleteUserByUUID(ctx *fiber.Ctx) error {
	var err error
	var userUUID uuid.UUID
	if userUUID, err = uuid.Parse(ctx.Query("uuid")); err != nil {
		log.Warnf("Error parsing user UUID: %v", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewHTTPError("cannot parse user UUID"))
	}

	timeoutCtx, cancel := context.WithTimeout(ctx.Context(), srv.cfg.MaxConnLifetime)
	defer cancel()

	conn, err := srv.db.Acquire(timeoutCtx)
	defer conn.Release()
	if err != nil {
		log.Warnf("A new database connection could not be established: %v", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.NewHTTPError("a new database connection could not be established"),
		)
	}

	if isExistsUser, err := srv.repo.CheckUserByUUID(&timeoutCtx, conn, &userUUID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewHTTPError("error checking the user's existence"))
	} else if !*isExistsUser {
		return ctx.Status(fiber.StatusNotFound).JSON(
			utils.NewHTTPError(fmt.Sprintf("user with uuid %v not found", userUUID)),
		)
	}
	if err = srv.repo.DeleteUserByUUID(&timeoutCtx, conn, &userUUID); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(
			utils.NewHTTPError(fmt.Sprintf("user with an UUID=%v cannot be deleted", userUUID)),
		)
	}
	return ctx.Status(fiber.StatusOK).JSON(schemes.HTTPResponse{Result: "success", Msg: "user deleted successfully"})
}
