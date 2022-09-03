package redissvc

import (
	"reflect"
	"strconv"
	"time"

	"github.com/ahl5esoft/lite-go/contract"
)

type cache struct {
	redis    contract.IRedis
	updateOn int64
	value    reflect.Value
	redisKey string
	loadFunc func() (interface{}, error)
}

func (m *cache) Flush() error {
	_, err := m.redis.HSet(
		"cache",
		m.redisKey,
		strconv.FormatInt(
			time.Now().UnixNano(),
			10,
		),
	)
	return err
}

func (m *cache) Get(k string, v interface{}) error {
	cacheOnStr, err := m.redis.HGet("cache", m.redisKey)
	if err != nil {
		return err
	}

	cacheOn, err := strconv.ParseInt(cacheOnStr, 10, 64)
	if err != nil {
		return err
	}

	if m.updateOn != cacheOn {
		m.updateOn = cacheOn
		res, err := m.loadFunc()
		if err != nil {
			return err
		}

		m.value = reflect.ValueOf(res)
	}

	rv := m.value.MapIndex(
		reflect.ValueOf(k),
	)
	if rv.IsValid() {
		reflect.ValueOf(v).Elem().Set(rv)
	}
	return nil
}

func (m *cache) GetUpdateOn() int64 {
	return m.updateOn
}

// 创建redis缓存
func NewCache(
	redis contract.IRedis,
	redisKey string,
	loadFunc func() (interface{}, error),
) contract.ICache {
	return &cache{
		redis:    redis,
		redisKey: redisKey,
		loadFunc: loadFunc,
		updateOn: time.Now().UnixNano(),
	}
}
