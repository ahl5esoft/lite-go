package contract

// 当前时间
type INowTime interface {
	Unix() int64
}
