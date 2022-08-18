package mongosvc

import (
	"context"
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/service/dbsvc"
)

type dbFactory struct {
	pool *dbPool
}

func (m *dbFactory) Db(entry mcontract.IDbModel, extra ...interface{}) contract.IDbRepository {
	var uow *unitOfWork
	isTx := true
	underscore.Chain(extra).Each(func(r interface{}, _ int) {
		if v, ok := r.(*unitOfWork); ok {
			uow = v
		}
	})

	if uow == nil {
		isTx = false
		uow = m.Uow().(*unitOfWork)
	}

	model := getModelMetadata(
		reflect.TypeOf(entry),
	)
	return dbsvc.NewDbRepository(uow, isTx, func() contract.IDbQuery {
		return newDbQuery(m.pool, model)
	})
}

func (m *dbFactory) Uow() contract.IUnitOfWork {
	return newUnitOfWork(m.pool)
}

func (m *dbFactory) WithContext(ctx context.Context) reflect.Value {
	return reflect.ValueOf(&dbFactory{
		pool: m.pool.WithContext(ctx),
	})
}

// 创建数据库工厂
func NewDbFactory(
	name string,
	uri string,
) contract.IDbFactory {
	return &dbFactory{
		pool: newDbPool(name, uri),
	}
}
