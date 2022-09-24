package contract

import (
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/model/global"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_UserValueServiceBase_GetNow(t *testing.T) {
	t.Run("global.UserValue", func(t *testing.T) {
		self := UserValueServiceBase{
			ValueServiceBase: ValueServiceBase[global.UserValue]{
				entries: []global.UserValue{
					{
						Value: map[int]int64{
							1: 99,
						},
					},
				},
			},
			NowValueType: 1,
		}

		res, err := self.GetNow()
		assert.Nil(t, err)
		assert.Equal(
			t,
			res,
			int64(99),
		)

		time.Sleep(time.Second)

		res, _ = self.GetNow()
		assert.Equal(
			t,
			res,
			int64(100),
		)
	})

	t.Run("NowTime", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockNowTime := NewMockINowTime(ctrl)
		self := UserValueServiceBase{
			ValueServiceBase: ValueServiceBase[global.UserValue]{
				entries: []global.UserValue{},
			},
			NowTime:      mockNowTime,
			NowValueType: 1,
		}

		mockNowTime.EXPECT().Unix().Return(
			int64(99),
		)

		res, err := self.GetNow()
		assert.Nil(t, err)
		assert.Equal(
			t,
			res,
			int64(99),
		)

		time.Sleep(time.Second)

		res, _ = self.GetNow()
		assert.Equal(
			t,
			res,
			int64(100),
		)
	})
}
