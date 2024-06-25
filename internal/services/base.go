package services

import (
	"github.com/CatLecter/gin_template/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type UserServiceInterface interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUserByUUID(ctx *fiber.Ctx) error
	UpdateUserByUUID(ctx *fiber.Ctx) error
	DeleteUserByUUID(ctx *fiber.Ctx) error
}

type Service struct{ UserServiceInterface }

func NewService(repos *repositories.Repository) *Service {
	return &Service{UserServiceInterface: NewUserService(repos.UserRepositoryInterface)}
}
