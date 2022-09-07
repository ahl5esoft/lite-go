package contract

// IApiFactory is api工厂
type IApiFactory interface {
	// 创建api
	Build(endpoint, api string) (instance IApi, err error)
	// 注册
	Register(endpoint, name string, api any)
}
