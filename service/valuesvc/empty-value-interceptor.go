package valuesvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/model/message"
)

type emptyValueInterceptor struct{}

func (m emptyValueInterceptor) After(_ contract.IUnitOfWork, _ contract.IValueService, _ message.ChangeValue) error {
	return nil
}

func (m emptyValueInterceptor) Before(_ contract.IUnitOfWork, _ contract.IValueService, _ message.ChangeValue) (bool, error) {
	return false, nil
}
