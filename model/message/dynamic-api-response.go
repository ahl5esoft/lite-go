package message

import errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"

// 动态api响应结构
type DynamicApiResponse[T any] struct {
	Data  T               `json:"data"`
	Error errorcode.Value `json:"err"`
}
