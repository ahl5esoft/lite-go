package contract

// 数值
type IValue interface {
	GetCount() int64
	GetSource() string
	GetValueType() int
}
