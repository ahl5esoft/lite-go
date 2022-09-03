package contract

import "github.com/ahl5esoft/lite-go/model/contract"

// 枚举
type IEnum interface {
	// 获取所有项
	GetAllItem() (map[int]contract.IEnumItem, error)
	// 获取所有项
	FindItems() ([]contract.IEnumItem, error)
}
