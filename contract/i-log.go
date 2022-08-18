package contract

// 日志
type ILog interface {
	AddLabel(k, f string, v ...interface{}) ILog
	Debug()
	Error(err error)
	Info()
	Warning()
}
