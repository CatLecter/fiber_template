package repositories

import (
	"github.com/CatLecter/gin_template/internal/schemes"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryInterface interface {
	CreateUser(user *schemes.UserRequest) (*schemes.UserResponse, error)
	GetUserByUUID(userUUID *uuid.UUID) (*schemes.UserResponse, error)
	UpdateUserByUUID(userUUID *uuid.UUID, user *schemes.UserRequest) (*schemes.UserResponse, error)
	DeleteUserByUUID(userUUID *uuid.UUID) error
	CheckUserByPhone(phone *string) (*bool, error)
}

type Repository struct{ UserRepositoryInterface }

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		UserRepositoryInterface: NewUserRepository(db),
	}
}
