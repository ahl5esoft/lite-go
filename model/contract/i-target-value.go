package contract

// 目标数值
type ITargetValue interface {
	IValue

	GetTargetNo() int
	GetTargetType() int
}
