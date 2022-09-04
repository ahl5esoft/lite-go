package contract

import "github.com/ahl5esoft/lite-go/model/contract"

// 枚举服务
type IEnumService[T contract.IEnumItem] interface {
	// 所有项
	AllItems() ([]T, error)
	// 获取项
	GetItem(func(T) bool) (T, error)
}
