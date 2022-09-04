package contract

// 数值日志
type IValueLog interface {
	IDbModel

	// 是否变更
	IsChange() bool
	// 设置数量
	SetCount(int64)
	// 设置主键
	SetID(string)
}
