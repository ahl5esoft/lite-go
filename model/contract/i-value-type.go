package contract

// 数值类型
type IValueType interface {
	IEnumItem

	// 获取每日数值
	GetDailyTime() int
	// 获取是否替换
	GetIsReplace() bool
}
