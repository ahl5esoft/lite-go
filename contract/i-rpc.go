package contract

// 远程过程调用
type IRpc interface {
	// 调用
	Call(string, interface{}) error
	// 调用(忽略响应错误)
	CallWithoutResponseError(string, interface{}) error
	// 设置请求体
	SetBody(any) IRpc
	// 设置头
	SetHeader(map[string]string) IRpc
}
