// Package repositories содержит репозитории для работы с данными пользователей.
package repositories

import (
	"context"
	"errors"
	"time"

	"fibertemplate/internal/schemes"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

// UserRepository реализует интерфейс для работы с пользователями в базе данных.
type UserRepository struct{}

// NewUserRepository создает новый экземпляр репозитория пользователей.
func NewUserRepository() UserRepositoryInterface { return &UserRepository{} }

// CreateUser создает нового пользователя в базе данных.
func (repo *UserRepository) CreateUser(ctx *context.Context, conn *pgxpool.Conn, user *schemes.UserRequest) (*schemes.UserResponse, error) {
	userResp := schemes.UserResponse{}
	query := "INSERT INTO users(full_name, phone) VALUES($1, $2) RETURNING *"
	if err := conn.QueryRow(*ctx, query, &user.FullName, &user.Phone).Scan(
		&userResp.UUID, &userResp.FullName, &userResp.Phone, &userResp.CreatedAt, &userResp.UpdatedAt,
	); err != nil {
		log.Warnf("Failed to insert user: %s", err.Error())
		return nil, err
	}
	return &userResp, nil
}

// GetUserByUUID получает пользователя по UUID из базы данных.
func (repo *UserRepository) GetUserByUUID(ctx *context.Context, conn *pgxpool.Conn, userUUID *uuid.UUID) (*schemes.UserResponse, error) {
	userResp := schemes.UserResponse{}
	query := "SELECT * FROM users WHERE uuid = $1"
	if err := conn.QueryRow(*ctx, query, &userUUID).Scan(
		&userResp.UUID, &userResp.FullName, &userResp.Phone, &userResp.CreatedAt, &userResp.UpdatedAt,
	); err != nil {
		log.Warnf("Failed to get user: %s", err.Error())
		return nil, err
	}
	return &userResp, nil
}

// CheckUserByUUID проверяет существование пользователя по UUID.
func (repo *UserRepository) CheckUserByUUID(ctx *context.Context, conn *pgxpool.Conn, userUUID *uuid.UUID) (*bool, error) {
	var result bool
	query := "SELECT CASE WHEN EXISTS (SELECT uuid FROM users WHERE uuid = $1) THEN true ELSE false END"
	if err := conn.QueryRow(*ctx, query, &userUUID).Scan(&result); err != nil {
		log.Warnf("Failed to get user: %s", err.Error())
		return nil, err
	}
	return &result, nil
}

// CheckUserByPhone проверяет существование пользователя по номеру телефона.
func (repo *UserRepository) CheckUserByPhone(ctx *context.Context, conn *pgxpool.Conn, phone *string) (*bool, error) {
	var result bool
	query := "SELECT CASE WHEN EXISTS (SELECT uuid FROM users WHERE phone = $1) THEN true ELSE false END"
	if err := conn.QueryRow(*ctx, query, &phone).Scan(&result); err != nil {
		log.Warnf("Failed to check user: %s", err.Error())
		return nil, err
	}
	return &result, nil
}

// UpdateUserByUUID обновляет данные пользователя по UUID.
func (repo *UserRepository) UpdateUserByUUID(ctx *context.Context, conn *pgxpool.Conn, userUUID *uuid.UUID, user *schemes.UserRequest) (*schemes.UserResponse, error) {
	userResp := schemes.UserResponse{}
	query := "UPDATE users SET full_name = $1, phone = $2, updated_at = $3 WHERE uuid = $4 RETURNING *"
	if err := conn.QueryRow(*ctx, query, user.FullName, user.Phone, time.Now(), userUUID).Scan(
		&userResp.UUID, &userResp.FullName, &userResp.Phone, &userResp.CreatedAt, &userResp.UpdatedAt,
	); err != nil {
		log.Warnf("Failed to update user: %s", err.Error())
		return nil, err
	}
	return &userResp, nil
}

// DeleteUserByUUID удаляет пользователя по UUID из базы данных.
func (repo *UserRepository) DeleteUserByUUID(ctx *context.Context, conn *pgxpool.Conn, userUUID *uuid.UUID) error {
	result, err := conn.Exec(*ctx, "DELETE FROM users WHERE uuid = $1 RETURNING TRUE", &userUUID)
	if err != nil {
		log.Warnf("Failed to get user: %s", err.Error())
		return err
	}
	if result.String() == "DELETE 0" {
		log.Warnf("User with UUID=%v does not exist", userUUID)
		return errors.New("user does not exist")
	}
	return nil
}
