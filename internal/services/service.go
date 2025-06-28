// Package services реализует бизнес-логику приложения.
package services

import (
	cfg "fibertemplate/internal/config"
	"fibertemplate/internal/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Service содержит все сервисы приложения.
type Service struct{ UserServiceInterface }

// NewService создает новый экземпляр сервиса.
func NewService(cfg *cfg.Config, db *pgxpool.Pool) *Service {
	repos := repositories.NewRepository()
	return &Service{UserServiceInterface: NewUserService(cfg, db, repos.UserRepositoryInterface)}
}
