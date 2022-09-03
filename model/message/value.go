package message

// 数值
type Value struct {
	// 数值类型
	ValueType int `validate:"min=0,required"`
	// 数量,
	Count int64 `validate:"required"`
	// 来源
	Source string
}

// 获取数量
func (m Value) GetCount() int64 {
	return m.Count
}

// 获取来源
func (m Value) GetSource() string {
	return m.Source
}

// 获取数值类型
func (m Value) GetValueType() int {
	return m.ValueType
}
