// Package repositories содержит репозитории для работы с данными.
package repositories

import (
	"context"

	"fibertemplate/internal/schemes"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserRepositoryInterface определяет интерфейс для работы с пользователями.
type UserRepositoryInterface interface {
	CreateUser(ctx *context.Context, conn *pgxpool.Conn, user *schemes.UserRequest) (*schemes.UserResponse, error)
	GetUserByUUID(ctx *context.Context, conn *pgxpool.Conn, userUUID *uuid.UUID) (*schemes.UserResponse, error)
	UpdateUserByUUID(ctx *context.Context, conn *pgxpool.Conn, userUUID *uuid.UUID, user *schemes.UserRequest) (*schemes.UserResponse, error)
	DeleteUserByUUID(ctx *context.Context, conn *pgxpool.Conn, userUUID *uuid.UUID) error
	CheckUserByUUID(ctx *context.Context, conn *pgxpool.Conn, userUUID *uuid.UUID) (*bool, error)
	CheckUserByPhone(ctx *context.Context, conn *pgxpool.Conn, phone *string) (*bool, error)
}

// Repository содержит все репозитории приложения.
type Repository struct{ UserRepositoryInterface }

// NewRepository создает новый экземпляр репозитория.
func NewRepository() *Repository {
	return &Repository{
		UserRepositoryInterface: NewUserRepository(),
	}
}
