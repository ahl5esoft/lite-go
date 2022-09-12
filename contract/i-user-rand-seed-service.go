package contract

// 用户随机种子服务
type IUserRandSeedService interface {
	// 获取
	Get(IUnitOfWork, int, int) (int, error)
	// 使用
	Use(IUnitOfWork, int) (int, error)
}
