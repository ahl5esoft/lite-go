package contract

import (
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/global"
)

// 目标数值服务
type ITargetValueService interface {
	// 验证条件
	CheckConditions(IUnitOfWork, [][]mcontract.IValueCondition) (bool, error)
	// 获取数量
	GetCount(IUnitOfWork, int) (int64, error)
	// 获取数据
	GetEntry(*global.UserValue) error
	// 更新
	Update(IUnitOfWork, string, []mcontract.IValue) error
}
