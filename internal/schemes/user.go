// Package schemes содержит структуры для пользователей.
package schemes

import "time"

// UserRequest представляет запрос на создание/обновление пользователя.
type UserRequest struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

// UserResponse представляет ответ с данными пользователя.
type UserResponse struct {
	UUID      string    `json:"uuid"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
