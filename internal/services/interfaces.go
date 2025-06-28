// Package services содержит интерфейсы для бизнес-логики приложения.
package services

import (
	"context"

	"fibertemplate/internal/schemes"
	"github.com/google/uuid"
)

// UserServiceInterface определяет интерфейс для работы с пользователями.
type UserServiceInterface interface {
	CreateUser(ctx context.Context, req *schemes.UserRequest) (*schemes.UserResponse, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*schemes.UserResponse, error)
	UpdateUserByID(ctx context.Context, id uuid.UUID, req *schemes.UserRequest) (*schemes.UserResponse, error)
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
}
