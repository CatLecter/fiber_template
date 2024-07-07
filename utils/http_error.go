package utils

import "github.com/CatLecter/gin_template/internal/schemes"

func NewHTTPError(msg string) *schemes.HTTPResponse {
	return &schemes.HTTPResponse{Result: "error", Msg: msg}
}
