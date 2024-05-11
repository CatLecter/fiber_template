package repositories

import (
	"app/dto"
	"app/engines"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

func GetUserByID(userID *uuid.UUID) (*dto.ResponseUserDTO, error) {
	ctx := context.Background()
	dbEngine := engines.DBEngine{}
	conn, err := dbEngine.GetConn(&ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	responseUser := dto.ResponseUserDTO{}
	err = conn.QueryRow(ctx, "SELECT * FROM users WHERE user_id = $1", userID).Scan(
		&responseUser.UserID,
		&responseUser.FullName,
		&responseUser.Phone,
		&responseUser.CreatedAt,
		&responseUser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &responseUser, nil
}

func CreateUser(requestUser *dto.RequestUserDTO) (*dto.ResponseUserDTO, error) {
	ctx := context.Background()
	dbEngine := engines.DBEngine{}
	conn, err := dbEngine.GetConn(&ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	responseUser := dto.ResponseUserDTO{}
	err = conn.QueryRow(
		ctx, "INSERT INTO users(full_name, phone) VALUES($1, $2) RETURNING *",
		requestUser.FullName,
		requestUser.Phone,
	).Scan(
		&responseUser.UserID,
		&responseUser.FullName,
		&responseUser.Phone,
		&responseUser.CreatedAt,
		&responseUser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &responseUser, nil
}

func UpdateUserByID(userID *uuid.UUID, requestUser *dto.RequestUserDTO) (*dto.ResponseUserDTO, error) {
	ctx := context.Background()
	dbEngine := engines.DBEngine{}
	conn, err := dbEngine.GetConn(&ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	responseUser := dto.ResponseUserDTO{}
	err = conn.QueryRow(
		ctx, "UPDATE users SET full_name = $1, phone = $2, updated_at = $3 WHERE user_id = $4 RETURNING *",
		requestUser.FullName,
		requestUser.Phone,
		time.Now(),
		userID,
	).Scan(
		&responseUser.UserID,
		&responseUser.FullName,
		&responseUser.Phone,
		&responseUser.CreatedAt,
		&responseUser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &responseUser, nil
}

func DeleteUserByID(userID *uuid.UUID) error {
	ctx := context.Background()
	dbEngine := engines.DBEngine{}
	conn, err := dbEngine.GetConn(&ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	result, err := conn.Exec(ctx, "DELETE FROM users WHERE user_id = $1", userID)
	if err != nil {
		return err
	}
	if result.String() == "DELETE 0" {
		return errors.New("WARNING: user has not been deleted")
	}
	return nil
}
