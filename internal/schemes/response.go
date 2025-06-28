// Package schemes содержит структуры для ответов API.
package schemes

// HTTPResponse представляет стандартный HTTP-ответ API.
type HTTPResponse struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
}
