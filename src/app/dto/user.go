package dto

import "time"

type RequestUserDTO struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type ResponseUserDTO struct {
	UserID    string    `json:"user_id"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
