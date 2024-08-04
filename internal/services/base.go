package services

import (
	"github.com/CatLecter/gin_template/configs"
	"github.com/CatLecter/gin_template/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserServiceInterface interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUserByUUID(ctx *fiber.Ctx) error
	UpdateUserByUUID(ctx *fiber.Ctx) error
	DeleteUserByUUID(ctx *fiber.Ctx) error
}

type Service struct{ UserServiceInterface }

func NewService(cfg *configs.Config, db *pgxpool.Pool, repos *repositories.Repository) *Service {
	return &Service{UserServiceInterface: NewUserService(cfg, db, repos.UserRepositoryInterface)}
}
