package global

// 用户数值日志
type UserValueLog struct {
	ValueType int
	Count     int64
	OldCount  int64
	ID        string `alias:"" bson:"_id" db:"_id"`
	Source    string
	UserID    string
}

func (m UserValueLog) GetID() string {
	return m.ID
}

func (m UserValueLog) IsChange() bool {
	return m.Count != m.OldCount
}

func (m *UserValueLog) SetCount(v int64) {
	m.Count = v
}

func (m *UserValueLog) SetID(v string) {
	m.ID = v
}
