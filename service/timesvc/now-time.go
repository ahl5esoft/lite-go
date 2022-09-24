package timesvc

import (
	"time"

	"github.com/ahl5esoft/lite-go/contract"
)

// 创建当前时间
func NewNowTime() contract.INowTime {
	return contract.NowTimeBase{
		GetUnixFunc: func() int64 {
			return time.Now().Unix()
		},
	}
}
