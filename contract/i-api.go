package contract

// IApi is api接口
type IApi interface {
	// Call is 调用
	Call() (any, error)
}
