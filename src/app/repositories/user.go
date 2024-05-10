package repositories

import (
	"app/engines"
	"app/schemes"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

func GetUserByUUID(userUUID *uuid.UUID) (*schemes.UserResp, error) {
	ctx := context.Background()
	dbEngine := engines.DBEngine{}
	conn, err := dbEngine.GetConn(&ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	user := schemes.UserResp{}
	err = conn.QueryRow(ctx, "SELECT * FROM users WHERE uuid = $1", userUUID).Scan(
		&user.UUID,
		&user.FullName,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *schemes.UserReq) (*schemes.UserResp, error) {
	ctx := context.Background()
	dbEngine := engines.DBEngine{}
	conn, err := dbEngine.GetConn(&ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	userResp := schemes.UserResp{}
	err = conn.QueryRow(
		ctx, "INSERT INTO users(full_name, phone) VALUES($1, $2) RETURNING *", user.FullName, user.Phone,
	).Scan(
		&userResp.UUID,
		&userResp.FullName,
		&userResp.Phone,
		&userResp.CreatedAt,
		&userResp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &userResp, nil
}

func UpdateUserByUUID(userUUID *uuid.UUID, user *schemes.UserReq) (*schemes.UserResp, error) {
	ctx := context.Background()
	dbEngine := engines.DBEngine{}
	conn, err := dbEngine.GetConn(&ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	userResp := schemes.UserResp{}
	err = conn.QueryRow(
		ctx, "UPDATE users SET full_name = $1, phone = $2, updated_at = $3 WHERE uuid = $4 RETURNING *",
		user.FullName,
		user.Phone,
		time.Now(),
		userUUID,
	).Scan(
		&userResp.UUID,
		&userResp.FullName,
		&userResp.Phone,
		&userResp.CreatedAt,
		&userResp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &userResp, nil
}

func DeleteUserByUUID(userUUID *uuid.UUID) error {
	ctx := context.Background()
	dbEngine := engines.DBEngine{}
	conn, err := dbEngine.GetConn(&ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	result, err := conn.Exec(ctx, "DELETE FROM users WHERE uuid = $1", userUUID)
	if err != nil {
		return err
	}
	if result.String() == "DELETE 0" {
		return errors.New("WARNING: user has not been deleted")
	}
	return nil
}
