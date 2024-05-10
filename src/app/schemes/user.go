package schemes

import "time"

type UserReq struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type UserResp struct {
	UUID      string    `json:"uuid"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
