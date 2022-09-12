package mongosvc

import (
	"container/list"
	"strconv"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/message"
)

type userRewardService[T mcontract.IValueType] struct {
	enumFactory contract.IEnumFactory
	userService contract.IUserService
}

func (m *userRewardService[T]) FindResults(uow contract.IUnitOfWork, rewards [][]message.Reward, scene, source string) (res []message.ChangeValue, err error) {
	if len(rewards) == 0 {
		return
	}

	var valueTypeEnum contract.IEnumService[T]
	if err = m.enumFactory.Build("ValueType", &valueTypeEnum); err != nil {
		return
	}

	var valueTypeItems []T
	if valueTypeItems, err = valueTypeEnum.AllItems(); err != nil {
		return
	}

	var valueTypeOfItem map[int]T
	underscore.Chain(valueTypeItems).Map(func(r T, _ int) []any {
		return []any{
			r.GetValue(),
			r,
		}
	}).Object().Value(valueTypeEnum)

	res = make([]message.ChangeValue, 0)
	for _, r := range rewards {
		var reward message.Reward
		if len(r) == 1 {
			reward = r[0]
		} else {
			total := 0
			for _, cr := range r {
				total += cr.Weight
			}

			var seed int
			seed, err = m.userService.GetRandSeedService(scene).Use(
				uow,
				len(
					strconv.Itoa(total),
				),
			)
			if err != nil {
				return
			}

			seed = seed % total
			for _, cr := range r {
				if seed -= cr.Weight; seed <= 0 {
					reward = cr
					break
				}
			}
		}

		if item, ok := valueTypeOfItem[reward.ValueType]; ok && item.GetRewards() != nil {
			for i := 0; i < int(reward.Count); i++ {
				var subValues []message.ChangeValue
				subValues, err = m.FindResults(
					uow,
					item.GetRewards().([][]message.Reward),
					scene,
					source,
				)
				if err != nil {
					return
				}

				res = append(res, subValues...)
			}
		} else {
			if reward.Source == "" {
				reward.Source = source
			}
			res = append(res, message.ChangeValue{
				Count:      reward.Count,
				Source:     reward.Source,
				TargetNo:   reward.TargetNo,
				TargetType: reward.TargetType,
				ValueType:  reward.ValueType,
			})
		}
	}
	return
}

func (m *userRewardService[T]) Preview(uow contract.IUnitOfWork, rewards [][]message.Reward, scene string) (res []message.ChangeValue, err error) {
	if len(rewards) == 0 {
		return
	}

	var valueTypeEnum contract.IEnumService[T]
	if err = m.enumFactory.Build("ValueType", &valueTypeEnum); err != nil {
		return
	}

	var valueTypeItems []T
	if valueTypeItems, err = valueTypeEnum.AllItems(); err != nil {
		return
	}

	var valueTypeOfItem map[int]T
	underscore.Chain(valueTypeItems).Map(func(r T, _ int) []any {
		return []any{
			r.GetValue(),
			r,
		}
	}).Object().Value(valueTypeEnum)

	res = make([]message.ChangeValue, 0)
	queue := list.New()
	for _, r := range rewards {
		queue.PushBack(r)
	}
	offset := 0
	for queue.Len() > 0 {
		group := queue.Remove(
			queue.Back(),
		).([]message.Reward)
		var reward message.Reward
		if len(group) == 1 {
			reward = group[0]
		} else {
			total := 0
			for _, cr := range group {
				total += cr.Weight
			}

			var seed int
			seed, err = m.userService.GetRandSeedService(scene).Get(
				uow,
				len(
					strconv.Itoa(total),
				),
				offset,
			)
			if err != nil {
				return
			}

			seed = seed % total
			for _, cr := range group {
				if seed -= cr.Weight; seed <= 0 {
					reward = cr
					break
				}
			}
		}

		if item, ok := valueTypeOfItem[reward.ValueType]; ok && item.GetRewards() != nil {
			for i := 0; i < int(reward.Count); i++ {
				for _, cr := range item.GetRewards().([][]message.Reward) {
					subQueue := list.New()
					for _, sr := range cr {
						subQueue.PushBack(sr)
					}
					queue.PushFrontList(subQueue)
				}
			}
		} else {
			res = append(res, message.ChangeValue{
				Count:     reward.Count,
				ValueType: reward.ValueType,
			})
		}
	}
	return
}

// 创建用户奖励服务
func NewUserRewardService[T mcontract.IValueType](
	enumFactory contract.IEnumFactory,
	userService contract.IUserService,
) contract.IUserRewardService {
	return &userRewardService[T]{
		enumFactory: enumFactory,
		userService: userService,
	}
}
