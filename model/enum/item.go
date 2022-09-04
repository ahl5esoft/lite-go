package enum

import "fmt"

// 枚举项
type Item struct {
	Value    int
	EnumName string
}

// 获取自定义多语言键
func (m Item) GetCustomEncodingKey(attr string) string {
	return fmt.Sprintf(
		"%s-%d-%s",
		m.EnumName,
		m.Value,
		attr,
	)
}

// 获取多语言键
func (m Item) GetEncodingKey() string {
	return m.GetCustomEncodingKey("name")
}

// 获取值
func (m Item) GetValue() int {
	return m.Value
}
