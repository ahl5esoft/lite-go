package rpcsvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/model/global"
	"github.com/ahl5esoft/lite-go/model/message"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_NewValueService_GetCount(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRpcFactory := contract.NewMockIRpcFactory(ctrl)
		self := NewValueService[global.UserValue](mockRpcFactory, "test", map[string]any{}, nil)

		mockRpc := contract.NewMockIRpc(ctrl)
		mockRpcFactory.EXPECT().Build().Return(mockRpc)

		mockRpc.EXPECT().SetBody(map[string]any{}).Return(mockRpc)

		mockRpc.EXPECT().Call(
			"/test/get-value",
			gomock.Any(),
		).SetArg(1, message.DynamicApiResponse[global.UserValue]{
			Data: global.UserValue{
				Value: map[int]int64{
					1: 11,
				},
			},
		})

		res, err := self.GetCount(nil, 1)
		assert.Nil(t, err)
		assert.Equal(
			t,
			res,
			int64(11),
		)
	})
}

func Test_NewValueService_Update(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRpcFactory := contract.NewMockIRpcFactory(ctrl)
		self := NewValueService[global.UserValue](mockRpcFactory, "test", nil, map[string]any{
			"a": "aa",
		})

		mockRpc := contract.NewMockIRpc(ctrl)
		mockRpcFactory.EXPECT().Build().Return(mockRpc)

		mockRpc.EXPECT().SetBody(nil).Return(mockRpc)

		mockRpc.EXPECT().Call(
			"/test/get-value",
			gomock.Any(),
		).SetArg(1, message.DynamicApiResponse[global.UserValue]{
			Data: global.UserValue{},
		})

		mockRpcFactory.EXPECT().Build().Return(mockRpc)

		mockRpc.EXPECT().SetBody(map[string]any{
			"changeValues": []message.ChangeValue{
				{
					Count:     1,
					ValueType: 11,
				},
			},
			"source": "test",
			"a":      "aa",
		}).Return(mockRpc)

		mockRpc.EXPECT().Call(
			"/test/update-values",
			gomock.Any(),
		).SetArg(1, message.DynamicApiResponse[any]{})

		err := self.Update(nil, "test", []message.ChangeValue{
			{
				Count:     1,
				ValueType: 11,
			},
		})
		assert.Nil(t, err)
	})
}
