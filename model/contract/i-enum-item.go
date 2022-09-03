package contract

// 枚举项
type IEnumItem interface {
	// 获取自定义多语言键
	GetCustomEncodingKey(string) string
	// 获取多语言键
	GetEncodingKey() string
	// 获取枚举值
	GetValue() int
}
