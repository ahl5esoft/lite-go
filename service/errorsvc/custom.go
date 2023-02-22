package errorsvc

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/contract"
	errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"
	jsoniter "github.com/json-iterator/go"
)

type custom struct {
	error

	code errorcode.Value
	data any
}

func (m custom) Error() string {
	s, _ := jsoniter.MarshalToString(map[string]any{
		"data": m.data,
		"err":  m.code,
	})
	return s
}

func (m custom) GetCode() errorcode.Value {
	return m.code
}

func (m custom) GetData() any {
	return m.data
}

// 创建自定义错误
func New(code errorcode.Value, data any) contract.IError {
	return custom{
		error: fmt.Errorf("%v", data),
		code:  code,
		data:  data,
	}
}

// 创建自定义错误
func Newf(code errorcode.Value, format string, args ...interface{}) contract.IError {
	return custom{
		error: fmt.Errorf(format, args...),
		code:  code,
		data:  fmt.Sprintf(format, args...),
	}
}
