package global

// 枚举模型
type Enum[T any] struct {
	ID    string `alias:"Enum" bson:"_id" db:"_id"`
	Items []T
}

func (m Enum[T]) GetID() string {
	return m.ID
}
