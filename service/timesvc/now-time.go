package timesvc

import (
	"time"

	"github.com/ahl5esoft/lite-go/contract"
)

type nowTime struct{}

func (m nowTime) Unix() int64 {
	return time.Now().Unix()
}

func NewNowTime() contract.INowTime {
	return new(nowTime)
}
