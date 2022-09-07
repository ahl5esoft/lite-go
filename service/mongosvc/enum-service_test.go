package mongosvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/model/enum"
	"github.com/ahl5esoft/lite-go/model/global"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_enumService_AllItems(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbFactory := contract.NewMockIDbFactory(ctrl)
		self := NewEnumService[enum.Item](mockDbFactory, "test")

		mockDbRepo := contract.NewMockIDbRepository(ctrl)
		mockDbFactory.EXPECT().Db(global.Enum[enum.Item]{}).Return(mockDbRepo)

		mockDbQuery := contract.NewMockIDbQuery(ctrl)
		mockDbRepo.EXPECT().Query().Return(mockDbQuery)

		mockDbQuery.EXPECT().Where(bson.M{
			"_id": "test",
		}).Return(mockDbQuery)

		mockDbQuery.EXPECT().ToArray(
			gomock.Any(),
		).SetArg(0, []global.Enum[enum.Item]{
			{
				Items: []enum.Item{
					{
						Value: 1,
					},
				},
			},
		})

		res, err := self.AllItems()
		assert.Nil(t, err)
		assert.EqualValues(t, res, []enum.Item{
			{
				Value: 1,
			},
		})
	})
}
