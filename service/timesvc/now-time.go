package timesvc

import (
	"time"

	"github.com/ahl5esoft/lite-go/contract"
)

type nowTime struct{}

func (m nowTime) IsSame(otherUnix int64, t string) bool {
	nT := time.Unix(
		m.Unix(),
		0,
	)
	nY, nM, nD := nT.Date()
	oT := time.Unix(otherUnix, 0)
	oY, oM, oD := oT.Date()
	switch t {
	case "day":
		return nY == oY && nM == oM && nD == oD
	case "month":
		return nY == oY && nM == oM
	case "week":
		nY, nW := nT.ISOWeek()
		oY, oW := oT.ISOWeek()
		return nY == oY && nW == oW
	case "year":
		return nY == oY
	default:
		return false
	}
}

func (m nowTime) Unix() int64 {
	return time.Now().Unix()
}

func NewNowTime() contract.INowTime {
	return new(nowTime)
}
