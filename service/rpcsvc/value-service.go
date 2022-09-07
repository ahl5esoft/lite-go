package rpcsvc

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/message"
)

// 创建rpc数值服务
func NewValueService[T mcontract.IValue](
	rpcFactory contract.IRpcFactory,
	app string,
	rpcGetBody map[string]any,
	rpcUpdateBody map[string]any,
) contract.IValueService {
	return &contract.ValueServiceBase[T]{
		FindEntriesFunc: func() (res []T, err error) {
			var resp message.DynamicApiResponse[T]
			err = rpcFactory.Build().SetBody(rpcGetBody).Call(
				fmt.Sprintf(
					"/%s/get-value",
					app,
				),
				&resp,
			)
			if err != nil {
				return
			}

			res = []T{resp.Data}
			return
		},
		UpdateFunc: func(uow contract.IUnitOfWork, source string, changeValues []message.ChangeValue, _ *[]T) error {
			var resp message.DynamicApiResponse[any]
			rpcBody := map[string]any{
				"changeValues": changeValues,
				"source":       source,
			}
			for k, v := range rpcUpdateBody {
				rpcBody[k] = v
			}
			return rpcFactory.Build().SetBody(rpcBody).Call(
				fmt.Sprintf(
					"/%s/update-values",
					app,
				),
				&resp,
			)
		},
	}
}
