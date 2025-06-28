// Package services реализует бизнес-логику приложения.
package services

import (
	cfg "fibertemplate/internal/config"
	"fibertemplate/internal/repositories"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserServiceInterface определяет интерфейс для работы с пользователями.
type UserServiceInterface interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUserByUUID(ctx *fiber.Ctx) error
	UpdateUserByUUID(ctx *fiber.Ctx) error
	DeleteUserByUUID(ctx *fiber.Ctx) error
}

// Service содержит все сервисы приложения.
type Service struct{ UserServiceInterface }

// NewService создает новый экземпляр сервиса.
func NewService(cfg *cfg.Config, db *pgxpool.Pool) *Service {
	repos := repositories.NewRepository()
	return &Service{UserServiceInterface: NewUserService(cfg, db, repos.UserRepositoryInterface)}
}
