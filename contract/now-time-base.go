package contract

import "time"

// 当前时间基类
type NowTimeBase struct {
	// 获取unix函数
	GetUnixFunc func() int64
}

func (m NowTimeBase) IsSame(otherUnix int64, t string) bool {
	nT := time.Unix(
		m.GetUnixFunc(),
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

// 当前unix
func (m NowTimeBase) Unix() int64 {
	return m.GetUnixFunc()
}
