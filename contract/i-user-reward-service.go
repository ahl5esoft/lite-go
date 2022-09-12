package contract

import "github.com/ahl5esoft/lite-go/model/message"

// 用户奖励服务
type IUserRewardService interface {
	// 获取结果
	FindResults(IUnitOfWork, [][]message.Reward, string, string) ([]message.ChangeValue, error)
	// 预览
	Preview(IUnitOfWork, [][]message.Reward, string) ([]message.ChangeValue, error)
}
