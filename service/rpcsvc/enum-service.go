package rpcsvc

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/message"
)

// 创建枚举服务
func NewEnumService[T mcontract.IEnumItem](
	rpcFactory contract.IRpcFactory,
	app, name string,
) contract.IEnumService[T] {
	return &contract.EnumServiceBase[T]{
		FindItemsFunc: func() (res []T, err error) {
			var resp message.DynamicApiResponse[[]T]
			err = rpcFactory.Build().SetBody(map[string]interface{}{
				"name": name,
			}).Call(
				fmt.Sprintf("/%s/find-enum-items", app),
				&resp,
			)
			if err != nil {
				return
			}

			res = resp.Data
			return
		},
	}
}
