package mongosvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
)

type UserServiceBase[T mcontract.IValueType] struct {
	DbFactory             contract.IDbFactory
	EnumFactory           contract.IEnumFactory
	UserID                string
	BuildValueServiceFunc func() contract.IValueService
	rewardService         contract.IUserRewardService
	valueService          contract.IValueService
	randSeedService       map[string]contract.IUserRandSeedService
}

func (m *UserServiceBase[T]) GetRandSeedService(scene string) contract.IUserRandSeedService {
	if _, ok := m.randSeedService[scene]; !ok {
		m.randSeedService[scene] = NewUserRandSeedService(
			m.DbFactory,
			scene,
			m.UserID,
		)
	}

	return m.randSeedService[scene]
}

func (m *UserServiceBase[T]) GetRewardService() contract.IUserRewardService {
	if m.rewardService == nil {
		m.rewardService = NewUserRewardService[T](m.EnumFactory, m)
	}

	return m.rewardService
}

func (m *UserServiceBase[T]) GetValueService() contract.IValueService {
	if m.valueService == nil {
		m.valueService = m.BuildValueServiceFunc()
	}

	return m.valueService
}
