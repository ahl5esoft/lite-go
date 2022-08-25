package logsvc

import "github.com/ahl5esoft/lite-go/contract"

type logFactory func() contract.ILog

func (m logFactory) Build() contract.ILog {
	return m()
}

// 创建日志工厂
func NewLogFactory(buildFunc func() contract.ILog) contract.ILogFactory {
	return logFactory(buildFunc)
}
