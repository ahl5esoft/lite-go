package mongosvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/global"
)

// 加载枚举项
func LoadEnumItems(dbFactory contract.IDbFactory) (res map[string][]mcontract.IEnumItem, err error) {
	var entries []global.Enum
	if err = dbFactory.Db(global.Enum{}).Query().ToArray(&entries); err != nil {
		return
	}

	return
}
