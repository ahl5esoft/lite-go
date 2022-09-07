package contract

// 配置加载器
type IConfigLoader interface {
	// 加载
	Load(v any) error
}
