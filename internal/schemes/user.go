package schemes

import "time"

type UserRequest struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type UserResponse struct {
	UUID      string    `json:"uuid"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
