package enumsvc

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/ahl5esoft/lite-go/contract"
)

type enumFactory struct {
	once       sync.Once
	enum       map[string]reflect.Value
	buildFuncs map[string]func() any
}

func (m *enumFactory) Build(name string, v any) (err error) {
	m.once.Do(func() {
		m.enum = map[string]reflect.Value{}
		for k, fn := range m.buildFuncs {
			m.enum[k] = reflect.ValueOf(
				fn(),
			)
		}
	})

	if cv, ok := m.enum[name]; ok {
		reflect.ValueOf(v).Elem().Set(cv)
	} else {
		err = fmt.Errorf("无效枚举: %s", name)
	}
	return
}

// 创建枚举工厂
func NewEnumFactory(buildFuncs map[string]func() any) contract.IEnumFactory {
	return &enumFactory{
		buildFuncs: buildFuncs,
	}
}
