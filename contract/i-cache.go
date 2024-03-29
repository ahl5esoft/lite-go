package contract

// 缓存
type ICache interface {
	// 刷新
	Flush() error
	// 获取
	Get(string, any) error
	// 更新时间
	GetUpdateOn() int64
}
