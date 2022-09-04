package contract

import (
	"github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/message"
)

// 数值服务
type IValueService interface {
	// 验证条件
	CheckConditions(IUnitOfWork, [][]contract.IValueCondition) (bool, error)
	// 获取数量
	GetCount(IUnitOfWork, int) (int64, error)
	// 获取数据
	GetEntry(any) error
	// 更新
	Update(IUnitOfWork, string, []message.ChangeValue) error
}
