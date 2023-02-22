package gosvc

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/contract"
	errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"
	jsoniter "github.com/json-iterator/go"
)

type customError struct {
	code errorcode.Value
	data any
}

func (m customError) Error() string {
	s, _ := jsoniter.MarshalToString(map[string]any{
		"data": m.data,
		"err":  m.code,
	})
	return s
}

func (m customError) GetCode() errorcode.Value {
	return m.code
}

func (m customError) GetData() any {
	return m.data
}

func New(code errorcode.Value, data any) contract.IError {
	return customError{
		code: code,
		data: data,
	}
}

func Newf(code errorcode.Value, format string, args ...any) contract.IError {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}
	return customError{
		code: code,
		data: format,
	}
}
