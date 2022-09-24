package contract

import (
	"time"

	"github.com/ahl5esoft/lite-go/model/global"
)

// 用户数值服务基类
type UserValueServiceBase struct {
	ValueServiceBase[global.UserValue]

	NowTime      INowTime
	NowValueType int

	now [2]int64
}

func (m *UserValueServiceBase) GetNow() (res int64, err error) {
	now := time.Now().Unix()
	if m.now[0] == 0 {
		var entry global.UserValue
		if err = m.GetEntry(&entry); err != nil {
			return
		}

		m.now = [2]int64{
			entry.Value[m.NowValueType],
			now,
		}
		if m.now[0] == 0 {
			m.now[0] = m.NowTime.Unix()
		}
	}

	res = m.now[0] + now - m.now[1]
	return
}
