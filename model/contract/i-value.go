package contract

// 目标数值
type IValue interface {
	IDbModel

	// 获取值
	GetValue() map[int]int64
}
