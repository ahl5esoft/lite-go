package genericsvc

type factory[T any] func() T

func (m factory[T]) Build() T {
	return m()
}

// 创建工厂
func NewFactory[TFactory any, TObject any](buildFunc func() TObject) TFactory {
	f := factory[TObject](buildFunc)
	return any(f).(TFactory)
}
