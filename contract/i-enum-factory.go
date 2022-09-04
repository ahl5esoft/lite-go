package contract

// 枚举工厂
type IEnumFactory interface {
	// 创建枚举
	Build(string, interface{}) error
}
