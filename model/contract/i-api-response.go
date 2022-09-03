package contract

import errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"

// api响应
type IApiResposne interface {
	// 获取数据
	GetData() interface{}
	// 获取错误码
	GetErrorCode() errorcode.Value
}
