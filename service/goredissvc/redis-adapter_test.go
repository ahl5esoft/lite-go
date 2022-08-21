package goredissvc

import (
	"context"
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

var self contract.IRedis
var ctx context.Context
var client redis.Cmdable

func Test_redisAdapter_Del(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		key := "Test_redisAdapter_Del"
		val := "test"
		client.Set(ctx, key, val, 0).Result()
		defer client.Del(ctx, key)

		res, err := self.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, res, val)

		count, err := self.Del(key)
		assert.NoError(t, err)
		assert.Equal(t, count, int64(1))

		res, err = self.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, res, "")
	})
}

func Test_redisAdapter_Get(t *testing.T) {
	t.Run("不存在", func(t *testing.T) {
		key := "Test_redisAdapter_Get_不存在"
		res, err := self.Get(key)
		assert.Nil(t, err)
		assert.Empty(t, res)
	})

	t.Run("ok", func(t *testing.T) {
		key := "Test_redisAdapter_Get_一致性"
		client.Set(ctx, key, "test", 0)

		result, err := self.Get(key)
		if err != nil || result != "test" {
			t.Error("err")
		}

		client.Del(ctx, key)
	})
}

func Test_redisAdapter_SetEXPX(t *testing.T) {
	t.Run("ex", func(t *testing.T) {
		k := "b"
		v := "test2"
		second := int64(60)
		res, err := self.Set(k, v, "ex", second)
		assert.NoError(t, err)
		assert.True(t, res)

		ttl, err := self.TTL(k)
		assert.NoError(t, err)
		assert.Equal(
			t,
			ttl,
			time.Duration(second*1000*1000*1000),
		)
	})

	t.Run("px", func(t *testing.T) {
		k := "c"
		v := "test3"
		millisecond := int32(10000)
		res, err := self.Set(k, v, "px", millisecond)
		assert.NoError(t, err)
		assert.True(t, res)

		ttl, err := self.TTL(k)
		assert.NoError(t, err)
		assert.Equal(
			t,
			ttl,
			time.Duration(millisecond)*time.Millisecond,
		)
	})
}

func Test_redisAdapter_SetXXNX(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		key := "a"
		val := "test1"
		set, err := self.Set(key, val, "xx")
		defer self.Del(key)
		assert.NoError(t, err)
		assert.Equal(t, set, false)

		res, err := self.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, res, "")

		set, err = self.Set(key, val, "nx")
		assert.NoError(t, err)
		assert.Equal(t, set, true)

		res, err = self.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, res, val)

		set, err = self.Set(key, val, "xx")
		assert.NoError(t, err)
		assert.Equal(t, set, true)

		res, err = self.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, res, val)
	})
}

func init() {
	opt := &redis.Options{
		Addr: "localhost:6379",
	}
	client = redis.NewClient(opt)
	self = NewRedis(
		OptionsRedisOption(opt),
	)
	ctx = context.Background()
}
