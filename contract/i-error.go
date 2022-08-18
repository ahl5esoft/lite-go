package contract

import errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"

// 错误接口
type IError interface {
	error

	// 代码
	GetCode() errorcode.Value
	// 数据
	GetData() any
}
