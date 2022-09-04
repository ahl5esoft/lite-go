package contract

// 当前时间
type INowTime interface {
	// 是否相同
	IsSame(int64, string) bool
	// 时间戳
	Unix() int64
}
