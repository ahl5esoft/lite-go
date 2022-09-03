package message

import errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"

// 动态api响应结构
type DynamicApiResponse[T any] struct {
	Data  T               `json:"data"`
	Error errorcode.Value `json:"err"`
}

// 获取数据
func (m DynamicApiResponse[T]) GetData() interface{} {
	return m.Data
}

// 获取错误码
func (m DynamicApiResponse[T]) GetErrorCode() errorcode.Value {
	return m.Error
}
