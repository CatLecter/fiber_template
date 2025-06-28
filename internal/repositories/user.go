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

// GetUserByID получает пользователя по ID из базы данных.
func (repo *UserRepository) GetUserByID(ctx *context.Context, conn *pgxpool.Conn, userID *uuid.UUID) (*schemes.UserResponse, error) {
	userResp := schemes.UserResponse{}
	query := "SELECT * FROM users WHERE uuid = $1"
	if err := conn.QueryRow(*ctx, query, &userID).Scan(
		&userResp.UUID, &userResp.FullName, &userResp.Phone, &userResp.CreatedAt, &userResp.UpdatedAt,
	); err != nil {
		log.Warnf("Failed to get user: %s", err.Error())
		return nil, err
	}
	return &userResp, nil
}

// CheckUserByID проверяет существование пользователя по ID.
func (repo *UserRepository) CheckUserByID(ctx *context.Context, conn *pgxpool.Conn, userID *uuid.UUID) (*bool, error) {
	var result bool
	query := "SELECT CASE WHEN EXISTS (SELECT uuid FROM users WHERE uuid = $1) THEN true ELSE false END"
	if err := conn.QueryRow(*ctx, query, &userID).Scan(&result); err != nil {
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

// UpdateUserByID обновляет данные пользователя по ID.
func (repo *UserRepository) UpdateUserByID(ctx *context.Context, conn *pgxpool.Conn, userID *uuid.UUID, user *schemes.UserRequest) (*schemes.UserResponse, error) {
	userResp := schemes.UserResponse{}
	query := "UPDATE users SET full_name = $1, phone = $2, updated_at = $3 WHERE uuid = $4 RETURNING *"
	if err := conn.QueryRow(*ctx, query, user.FullName, user.Phone, time.Now(), userID).Scan(
		&userResp.UUID, &userResp.FullName, &userResp.Phone, &userResp.CreatedAt, &userResp.UpdatedAt,
	); err != nil {
		log.Warnf("Failed to update user: %s", err.Error())
		return nil, err
	}
	return &userResp, nil
}

// DeleteUserByID удаляет пользователя по ID из базы данных.
func (repo *UserRepository) DeleteUserByID(ctx *context.Context, conn *pgxpool.Conn, userID *uuid.UUID) error {
	result, err := conn.Exec(*ctx, "DELETE FROM users WHERE uuid = $1 RETURNING TRUE", &userID)
	if err != nil {
		log.Warnf("Failed to delete user: %s", err.Error())
		return err
	}
	if result.RowsAffected() == 0 {
		log.Warnf("User with ID=%v does not exist", userID)
		return errors.New("user not found")
	}
	return nil
}
