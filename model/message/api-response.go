package message

import errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"

type ApiResponse struct {
	Data  any             `json:"data"`
	Error errorcode.Value `json:"err"`
}
