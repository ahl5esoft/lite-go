package contract

// 日志工厂
type ILogFactory interface {
	Build() ILog
}
