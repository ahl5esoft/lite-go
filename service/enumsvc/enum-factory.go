package enumsvc

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/contract"
)

type enumFactory map[string]func() contract.IEnum

func (m enumFactory) Build(name string) contract.IEnum {
	if v, ok := m[name]; ok {
		return v()
	}

	panic(
		fmt.Sprintf("无效枚举: %s", name),
	)
}

// 创建枚举工厂
func NewEnumFactory(buildFunc map[string]func() contract.IEnum) contract.IEnumFactory {
	return enumFactory(buildFunc)
}
