package message

// 奖励
type Reward struct {
	ChangeValue

	// 权重
	Weight int `validate:"min=1"`
}
