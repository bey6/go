package models

// Response 响应体
type Response struct {
	Code rune              `json:"code"`
	Msg  string            `json:"msg"`
	Data map[string]string `json:"data"`
}
