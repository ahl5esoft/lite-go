package message

// 变更数值
type ChangeValue struct {
	// 数值类型
	ValueType int `validate:"min=0,required"`
	// 目标编号 / 目标类型
	TargetNo, TargetType int `validate:"min=0"`
	// 数量,
	Count int64 `validate:"required"`
	// 来源
	Source string
}
