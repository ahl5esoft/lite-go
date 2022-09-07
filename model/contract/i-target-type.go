package contract

// 目标类型
type ITargetType interface {
	IEnumItem

	// 获取应用
	GetApp() string
}
