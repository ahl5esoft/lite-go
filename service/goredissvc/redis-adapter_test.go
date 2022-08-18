package goredissvc

import (
	"context"
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/stretchr/testify/assert"

	"github.com/go-redis/redis/v8"
)

var redisClient contract.IRedis
var ctx context.Context
var client redis.Cmdable

func Test_goRedis_Get(t *testing.T) {
	t.Run("不存在", func(t *testing.T) {
		key := "Test_goRedis_Get_不存在"
		res, err := redisClient.Get(key)
		assert.Nil(t, err)
		assert.Empty(t, res)
	})

	t.Run("ok", func(t *testing.T) {
		key := "Test_goRedis_Get_一致性"
		client.Set(ctx, key, "test", 0)

		result, err := redisClient.Get(key)
		if err != nil || result != "test" {
			t.Error("err")
		}

		client.Del(ctx, key)
	})
}

func init() {
	opt := &redis.Options{
		Addr: "localhost:6379",
	}
	client = redis.NewClient(opt)
	redisClient = NewRedis(
		RedisOptions(opt),
	)
	ctx = context.Background()
}
