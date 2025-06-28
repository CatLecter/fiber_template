// Package services реализует бизнес-логику, связанную с пользователями.
package services

import (
	"context"
	"fmt"

	cfg "fibertemplate/internal/config"
	"fibertemplate/internal/repositories"
	"fibertemplate/internal/schemes"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

// UserService реализует бизнес-логику для работы с пользователями.
type UserService struct {
	cfg  *cfg.Config
	db   *pgxpool.Pool
	repo repositories.UserRepositoryInterface
}

// NewUserService создает новый экземпляр сервиса пользователей.
func NewUserService(cfg *cfg.Config, db *pgxpool.Pool, repo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{cfg, db, repo}
}

// GetUserByID получает пользователя по ID.
func (srv *UserService) GetUserByID(ctx context.Context, userID uuid.UUID) (*schemes.UserResponse, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, srv.cfg.MaxConnLifetime)
	defer cancel()
	conn, err := srv.db.Acquire(timeoutCtx)
	if err != nil {
		log.Warnf("A new database connection could not be established: %v", err.Error())
		return nil, fmt.Errorf("a new database connection could not be established")
	}
	defer conn.Release()
	userResp, err := srv.repo.GetUserByID(&timeoutCtx, conn, &userID)
	if err != nil {
		return nil, fmt.Errorf("user with ID=%v not found", userID)
	}
	return userResp, nil
}

// CreateUser создает нового пользователя.
func (srv *UserService) CreateUser(ctx context.Context, newUser *schemes.UserRequest) (*schemes.UserResponse, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, srv.cfg.MaxConnLifetime)
	defer cancel()
	conn, err := srv.db.Acquire(timeoutCtx)
	if err != nil {
		log.Warnf("A new database connection could not be established: %v", err.Error())
		return nil, fmt.Errorf("a new database connection could not be established")
	}
	defer conn.Release()
	isExists, err := srv.repo.CheckUserByPhone(&timeoutCtx, conn, &newUser.Phone)
	if err != nil {
		return nil, fmt.Errorf("error checking the user's existence")
	}
	if *isExists {
		return nil, fmt.Errorf("user with phone number %v already exists", newUser.Phone)
	}
	userResp, err := srv.repo.CreateUser(&timeoutCtx, conn, newUser)
	if err != nil {
		log.Warnf("Error when creating a user: %v", err.Error())
		return nil, fmt.Errorf("user cannot be created")
	}
	return userResp, nil
}

// UpdateUserByID обновляет пользователя по ID.
func (srv *UserService) UpdateUserByID(ctx context.Context, userID uuid.UUID, user *schemes.UserRequest) (*schemes.UserResponse, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, srv.cfg.MaxConnLifetime)
	defer cancel()
	conn, err := srv.db.Acquire(timeoutCtx)
	if err != nil {
		log.Warnf("A new database connection could not be established: %v", err.Error())
		return nil, fmt.Errorf("a new database connection could not be established")
	}
	defer conn.Release()
	isExistsUser, err := srv.repo.CheckUserByID(&timeoutCtx, conn, &userID)
	if err != nil {
		return nil, fmt.Errorf("error checking the user's existence")
	}
	if !*isExistsUser {
		return nil, fmt.Errorf("user with id %v not found", userID)
	}
	userResp, err := srv.repo.UpdateUserByID(&timeoutCtx, conn, &userID, user)
	if err != nil {
		return nil, fmt.Errorf("user cannot be updated")
	}
	return userResp, nil
}

// DeleteUserByID удаляет пользователя по ID.
func (srv *UserService) DeleteUserByID(ctx context.Context, userID uuid.UUID) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, srv.cfg.MaxConnLifetime)
	defer cancel()
	conn, err := srv.db.Acquire(timeoutCtx)
	if err != nil {
		log.Warnf("A new database connection could not be established: %v", err.Error())
		return fmt.Errorf("a new database connection could not be established")
	}
	defer conn.Release()
	isExistsUser, err := srv.repo.CheckUserByID(&timeoutCtx, conn, &userID)
	if err != nil {
		return fmt.Errorf("error checking the user's existence")
	}
	if !*isExistsUser {
		return fmt.Errorf("user with id %v not found", userID)
	}
	if err = srv.repo.DeleteUserByID(&timeoutCtx, conn, &userID); err != nil {
		return fmt.Errorf("user with an ID=%v cannot be deleted", userID)
	}
	return nil
}
