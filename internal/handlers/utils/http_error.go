// Package handlerutilities содержит утилиты для обработки HTTP-ошибок в хендлерах.
package handlerutilities

import "fibertemplate/internal/schemes"

// NewHTTPError создает новый HTTP-ответ с ошибкой.
func NewHTTPError(msg string) *schemes.HTTPResponse {
	return &schemes.HTTPResponse{Result: "error", Msg: msg}
}
