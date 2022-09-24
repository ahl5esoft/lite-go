package contract

type UserServiceBase struct {
	DbFactory                IDbFactory
	EnumFactory              IEnumFactory
	UserID                   string
	BuildRewardServiceFunc   func(IEnumFactory, IUserService) IUserRewardService
	BuildRandSeedServiceFunc func(dbFactory IDbFactory, scene string, userID string) IUserRandSeedService
	BuildValueServiceFunc    func() IValueService

	rewardService   IUserRewardService
	valueService    IValueService
	randSeedService map[string]IUserRandSeedService
}

func (m *UserServiceBase) GetRandSeedService(scene string) IUserRandSeedService {
	if m.randSeedService == nil {
		m.randSeedService = make(map[string]IUserRandSeedService, 0)
	}

	if _, ok := m.randSeedService[scene]; !ok {
		m.randSeedService[scene] = m.BuildRandSeedServiceFunc(
			m.DbFactory,
			scene,
			m.UserID,
		)
	}

	return m.randSeedService[scene]
}

func (m *UserServiceBase) GetRewardService() IUserRewardService {
	if m.rewardService == nil {
		m.rewardService = m.BuildRewardServiceFunc(m.EnumFactory, m)
	}

	return m.rewardService
}

func (m *UserServiceBase) GetValueService() IValueService {
	if m.valueService == nil {
		m.valueService = m.BuildValueServiceFunc()
	}

	return m.valueService
}
