package logsvc

import "github.com/ahl5esoft/lite-go/contract"

type logFactory struct {
	buildFunc func() contract.ILog
}

func (m logFactory) Build() contract.ILog {
	return m.buildFunc()
}

// 创建日志工厂
func NewLogFactory(buildFunc func() contract.ILog) contract.ILogFactory {
	return &logFactory{
		buildFunc: buildFunc,
	}
}
