package mongosvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/global"
	"go.mongodb.org/mongo-driver/bson"
)

func NewEnumService[T mcontract.IEnumItem](
	dbFactory contract.IDbFactory,
	name string,
) contract.IEnumService[T] {
	return &contract.EnumServiceBase[T]{
		FindItemsFunc: func() (res []T, err error) {
			var entries []global.Enum[T]
			err = dbFactory.Db(global.Enum[T]{}).Query().Where(bson.M{
				"_id": name,
			}).ToArray(&entries)
			if err != nil {
				return
			}

			if len(entries) > 0 {
				res = entries[0].Items
			} else {
				res = make([]T, 0)
			}

			return
		},
	}
}
