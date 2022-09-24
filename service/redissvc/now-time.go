package redissvc

import (
	"time"

	"github.com/ahl5esoft/lite-go/contract"
)

// 创建当前时间
func NewNowTime(
	redis contract.IRedis,
) contract.INowTime {
	return contract.NowTimeBase{
		GetUnixFunc: func() int64 {
			if t, err := redis.Time(); err != nil {
				return time.Now().Unix()
			} else {
				return t.Unix()
			}
		},
	}
}
