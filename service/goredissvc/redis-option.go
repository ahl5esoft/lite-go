package goredissvc

import (
	"github.com/ahl5esoft/lite-go/service/redissvc"

	"github.com/go-redis/redis/v8"
)

// redis选项
func RedisOptions(options *redis.Options) redissvc.RedisOption[*redisAdapter] {
	return func(adapter *redisAdapter) {
		adapter.options = options
	}
}
