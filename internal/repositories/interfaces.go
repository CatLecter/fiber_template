// Package repositories содержит интерфейсы для работы с данными.
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
	GetUserByID(ctx *context.Context, conn *pgxpool.Conn, userID *uuid.UUID) (*schemes.UserResponse, error)
	UpdateUserByID(ctx *context.Context, conn *pgxpool.Conn, userID *uuid.UUID, user *schemes.UserRequest) (*schemes.UserResponse, error)
	DeleteUserByID(ctx *context.Context, conn *pgxpool.Conn, userID *uuid.UUID) error
	CheckUserByID(ctx *context.Context, conn *pgxpool.Conn, userID *uuid.UUID) (*bool, error)
	CheckUserByPhone(ctx *context.Context, conn *pgxpool.Conn, phone *string) (*bool, error)
}
