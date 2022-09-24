package contract

// 用户数值服务
type IUserValueService interface {
	IValueService

	// 获取当前时间
	GetNow() (int64, error)
}
