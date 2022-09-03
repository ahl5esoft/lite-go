package contract

// 数值条件
type IValueCondition interface {
	// 获取数量
	GetCount() int64
	// 获取操作符
	GetOp() string
	// 获取数值类型
	GetValueType() int
}
