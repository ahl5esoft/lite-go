package dbsvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
)

type dbRepository struct {
	uow             contract.IUnitOfWorkRepository
	isTx            bool
	createQueryFunc func() contract.IDbQuery
}

func (m dbRepository) Add(entry mcontract.IDbModel) error {
	m.uow.RegisterAdd(entry)

	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

func (m dbRepository) Query() contract.IDbQuery {
	return m.createQueryFunc()
}

func (m dbRepository) Remove(entry mcontract.IDbModel) error {
	m.uow.RegisterRemove(entry)

	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

func (m dbRepository) Save(entry mcontract.IDbModel) error {
	m.uow.RegisterSave(entry)

	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

// 创建数据仓储
func NewDbRepository(
	uow contract.IUnitOfWorkRepository,
	isTx bool,
	createQueryFunc func() contract.IDbQuery,
) contract.IDbRepository {
	return &dbRepository{
		createQueryFunc: createQueryFunc,
		isTx:            isTx,
		uow:             uow,
	}
}
