package contract

// 用户服务
type IUserService interface {
	// 获取随机种子服务
	GetRandSeedService(string) IUserRandSeedService
	// 获取随机种子服务
	GetRewardService() IUserRewardService
	// 获取数值服务
	GetValueService() IValueService
}
