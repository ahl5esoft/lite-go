package redissvc

import "github.com/ahl5esoft/lite-go/contract"

// redis选项
type RedisOption[T contract.IRedis] func(redis T)
