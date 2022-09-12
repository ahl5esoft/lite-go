package mongosvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/model/global"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_userRandSeedService_Get(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbFactory := contract.NewMockIDbFactory(ctrl)
		self := NewUserRandSeedService(mockDbFactory, "", "user-id")

		mockDbRepo := contract.NewMockIDbRepository(ctrl)
		mockDbFactory.EXPECT().Db(global.UserRandSeed{}, nil).Return(mockDbRepo)

		self.(*userRandSeedService).entries = []global.UserRandSeed{
			{
				ID: "",
				Seed: map[string]string{
					"": "563218",
				},
			},
		}

		res, err := self.Get(nil, 2, 0)
		assert.Nil(t, err)
		assert.Equal(t, res, 56)
	})

	t.Run("偏移", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbFactory := contract.NewMockIDbFactory(ctrl)
		self := NewUserRandSeedService(mockDbFactory, "", "user-id")

		mockDbRepo := contract.NewMockIDbRepository(ctrl)
		mockDbFactory.EXPECT().Db(global.UserRandSeed{}, nil).Return(mockDbRepo)

		self.(*userRandSeedService).entries = []global.UserRandSeed{
			{
				ID: "",
				Seed: map[string]string{
					"": "563218125",
				},
			},
		}

		res, err := self.Get(nil, 4, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, 6321)
	})
}

func Test_userRandSeedService_Use(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbFactory := contract.NewMockIDbFactory(ctrl)
		self := NewUserRandSeedService(mockDbFactory, "", "user-id")

		mockDbRepo := contract.NewMockIDbRepository(ctrl)
		mockDbFactory.EXPECT().Db(global.UserRandSeed{}, nil).Return(mockDbRepo)

		self.(*userRandSeedService).entries = []global.UserRandSeed{
			{
				ID: "",
				Seed: map[string]string{
					"": "563218125",
				},
			},
		}

		mockDbFactory.EXPECT().Db(global.UserRandSeed{}, nil).Return(mockDbRepo)

		mockDbRepo.EXPECT().Save(global.UserRandSeed{
			ID: "",
			Seed: map[string]string{
				"": "18125",
			},
		})

		res, err := self.Use(nil, 4)
		assert.Nil(t, err)
		assert.Equal(t, res, 5632)
	})
}

func Test_userRandSeedService_getSeed(t *testing.T) {
	t.Run("新增", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbFactory := contract.NewMockIDbFactory(ctrl)
		self := NewUserRandSeedService(mockDbFactory, "getSeed", "user-id")

		mockDbRepo := contract.NewMockIDbRepository(ctrl)
		mockDbFactory.EXPECT().Db(global.UserRandSeed{}, nil).Return(mockDbRepo)

		mockDbQuery := contract.NewMockIDbQuery(ctrl)
		mockDbRepo.EXPECT().Query().Return(mockDbQuery)

		mockDbQuery.EXPECT().Where(bson.M{
			"_id": "user-id",
		}).Return(mockDbQuery)

		mockDbQuery.EXPECT().ToArray(
			gomock.Any(),
		).SetArg(0, []global.UserRandSeed{})

		mockDbRepo.EXPECT().Add(global.UserRandSeed{
			ID:   "user-id",
			Seed: make(map[string]string),
		})

		mockDbRepo.EXPECT().Save(
			gomock.Any(),
		)

		res, err := self.(*userRandSeedService).getSeed(nil)
		assert.Nil(t, err)
		assert.True(
			t,
			len(res) >= 300,
		)
	})
}

func init() {
	UserRandSeedLengthRange[""] = [2]int{0, 50}
	UserRandSeedLengthRange["getSeed"] = [2]int{150, 300}
}
