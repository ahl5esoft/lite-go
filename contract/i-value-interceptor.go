package contract

import (
	"github.com/ahl5esoft/lite-go/model/message"
)

// 数值拦截器
type IValueInterceptor interface {
	// 后置
	After(IUnitOfWork, IValueService, message.ChangeValue) error
	// 前置
	Before(IUnitOfWork, IValueService, message.ChangeValue) (bool, error)
}
