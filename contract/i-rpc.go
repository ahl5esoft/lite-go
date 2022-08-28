package contract

import "github.com/ahl5esoft/lite-go/model/message"

// 远程过程调用
type IRpc interface {
	// 调用
	Call(string, *message.Api) error
	// 设置请求体
	SetBody(any) IRpc
	// 设置头
	SetHeader(map[string]string) IRpc
}
