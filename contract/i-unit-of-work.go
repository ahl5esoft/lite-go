package contract

// 工作单元
type IUnitOfWork interface {
	// 提交
	Commit() error
	// 注册提交后事件
	RegisterAfter(func() error, string)
}
