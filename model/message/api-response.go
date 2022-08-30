package message

import errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"

// api响应结构
type ApiResponse struct {
	Data  interface{}     `json:"data"`
	Error errorcode.Value `json:"err"`
}
