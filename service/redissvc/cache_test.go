package redissvc

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_cache_Flush(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRedis := contract.NewMockIRedis(ctrl)
		self := NewCache(mockRedis, "redis-key", nil)

		mockRedis.EXPECT().HSet(
			"cache",
			"redis-key",
			gomock.Len(
				len(
					strconv.FormatInt(
						time.Now().UnixNano(),
						10,
					),
				),
			),
		)

		self.Flush()
	})
}

func Test_cache_Get(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRedis := contract.NewMockIRedis(ctrl)
		self := NewCache(mockRedis, "redis-key", nil)

		mockRedis.EXPECT().HGet("cache", "redis-key").Return("1", nil)

		self.(*cache).updateOn = 1
		self.(*cache).value = reflect.ValueOf(map[string][]int{
			"": {1, 2, 3},
		})

		var res []int
		err := self.Get("", &res)
		assert.Nil(t, err)
		assert.EqualValues(t, res, []int{1, 2, 3})
	})

	t.Run("加载", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRedis := contract.NewMockIRedis(ctrl)
		loadCount := 0
		self := NewCache(mockRedis, "redis-key", func() (interface{}, error) {
			loadCount++
			return map[string][]int{
				"": {1, 2, 3},
			}, nil
		})

		mockRedis.EXPECT().HGet("cache", "redis-key").Return("11", nil)
		mockRedis.EXPECT().HGet("cache", "redis-key").Return("11", nil)

		var res []int
		err := self.Get("", &res)
		assert.Nil(t, err)
		assert.EqualValues(t, res, []int{1, 2, 3})

		self.(*cache).updateOn = 11
		assert.EqualValues(
			t,
			self.(*cache).value.Interface().(map[string][]int),
			map[string][]int{
				"": {1, 2, 3},
			},
		)

		self.Get("", &res)

		assert.Equal(t, loadCount, 1)
	})
}
