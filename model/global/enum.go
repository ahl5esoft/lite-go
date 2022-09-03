package global

// 枚举模型
type Enum struct {
	ID    string `alias:"" bson:"_id" db:"_id"`
	Items map[interface{}]interface{}
}

func (m Enum) GetID() string {
	return m.ID
}
