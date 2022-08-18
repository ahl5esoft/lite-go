package contract

// 配置加载器
type IConfigLoader interface {
	Load(v interface{}) error
}
