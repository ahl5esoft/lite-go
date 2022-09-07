package rpcsvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/model/enum"
	"github.com/ahl5esoft/lite-go/model/message"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_NewEnumService_GetItem(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRpcFactory := contract.NewMockIRpcFactory(ctrl)
		self := NewEnumService[enum.Item](mockRpcFactory, "app", "test")

		mockRpc := contract.NewMockIRpc(ctrl)
		mockRpcFactory.EXPECT().Build().Return(mockRpc)

		mockRpc.EXPECT().SetBody(map[string]interface{}{
			"name": "test",
		}).Return(mockRpc)

		mockRpc.EXPECT().Call(
			"/app/find-enum-items",
			gomock.Any(),
		).SetArg(1, message.DynamicApiResponse[[]enum.Item]{
			Data: []enum.Item{
				{
					Value: 1,
				},
			},
		})

		res, err := self.GetItem(func(r enum.Item) bool {
			return r.Value == 1
		})
		assert.Nil(t, err)
		assert.Equal(t, res, enum.Item{
			Value: 1,
		})
	})
}
