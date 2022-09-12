package message

// 奖励
type Reward struct {
	ChangeValue `bson:",inline"`

	// 权重
	Weight int `validate:"min=1"`
}
