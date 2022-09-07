package contract

import "github.com/ahl5esoft/lite-go/model/contract"

// 数据工厂
type IDbFactory interface {
	// 仓储
	Db(entry contract.IDbModel, extra ...any) IDbRepository
	// 工作单元
	Uow() IUnitOfWork
}
