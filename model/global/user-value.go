package global

// 用户数值
type UserValue struct {
	ID    string `alias:"" bson:"_id" db:"_id"`
	Value map[int]int64
}

func (m UserValue) GetID() string {
	return m.ID
}
