package contract

// rpc工厂
type IRpcFactory interface {
	// 创建rpc
	Build() IRpc
}
