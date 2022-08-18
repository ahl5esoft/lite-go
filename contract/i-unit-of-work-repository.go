package contract

import "github.com/ahl5esoft/lite-go/model/contract"

// 工作单元仓储
type IUnitOfWorkRepository interface {
	IUnitOfWork

	RegisterAdd(entry contract.IDbModel)
	RegisterSave(entry contract.IDbModel)
	RegisterRemove(entry contract.IDbModel)
}
