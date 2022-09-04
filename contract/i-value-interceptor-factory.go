package contract

import (
	"github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/message"
)

// 数值拦截器工厂
type IValueInterceptorFactory interface {
	// 创建拦截器
	Build(message.ChangeValue) (IValueInterceptor, error)
	// 注冊拦截器
	Register(int, IValueInterceptor)
	// 注冊断言拦截器
	RegisterPredicate(func(contract.IEnumItem) bool, IValueInterceptor)
}
