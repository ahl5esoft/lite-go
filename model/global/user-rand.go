package global

// 用户随机种子
type UserRandSeed struct {
	ID   string `alias:"" bson:"_id" db:"_id"`
	Seed map[string]string
}

func (m UserRandSeed) GetID() string {
	return m.ID
}
