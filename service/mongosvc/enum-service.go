package mongosvc

import (
	"sync"

	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/global"
	"go.mongodb.org/mongo-driver/bson"
)

type enumService[T mcontract.IEnumItem] struct {
	dbFactory contract.IDbFactory
	name      string
	once      sync.Once
	items     []T
}

func (m *enumService[T]) AllItems() (res []T, err error) {
	m.once.Do(func() {
		if m.items != nil {
			return
		}

		var entries []global.Enum[T]
		err = m.dbFactory.Db(global.Enum[T]{}).Query().Where(bson.M{
			"_id": m.name,
		}).ToArray(&entries)
		if err != nil {
			return
		}

		if len(entries) > 0 {
			m.items = entries[0].Items
		} else {
			m.items = make([]T, 0)
		}
	})

	if err == nil {
		res = m.items
	}

	return
}

func (m *enumService[T]) GetItem(predicate func(T) bool) (res T, err error) {
	var items []T
	if items, err = m.AllItems(); err != nil {
		return
	}

	for _, r := range items {
		if predicate(r) {
			res = r
			break
		}
	}

	return
}

func NewEnumService[T mcontract.IEnumItem](
	dbFactory contract.IDbFactory,
	name string,
) contract.IEnumService[T] {
	return &enumService[T]{
		dbFactory: dbFactory,
		name:      name,
	}
}
