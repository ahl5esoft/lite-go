package redissvc

import (
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_NewNowTime_Unix(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRedis := contract.NewMockIRedis(ctrl)
		self := NewNowTime(mockRedis)

		mockRedis.EXPECT().Time().Return(
			time.Unix(99, 0),
			nil,
		)

		res := self.Unix()
		assert.Equal(
			t,
			res,
			int64(99),
		)
	})
}
