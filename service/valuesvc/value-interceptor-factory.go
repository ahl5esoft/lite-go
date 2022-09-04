package valuesvc

import (
	"reflect"
	"sync"

	"github.com/ahl5esoft/lite-go/contract"
	mcontract "github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/message"
	"github.com/ahl5esoft/lite-go/service/iocsvc"
)

type predicateInterceptor struct {
	interceptor contract.IValueInterceptor
	predicate   func(mcontract.IEnumItem) bool
}

// 数值拦截器工厂
type valueInterceptorFactory[T mcontract.IEnumItem] struct {
	defaultInterceptor    contract.IValueInterceptor
	enumFacotry           contract.IEnumFactory
	once                  sync.Once
	metadata              map[int]reflect.Type
	predicateInterceptors []predicateInterceptor
	valueTypeInterceptor  map[int]contract.IValueInterceptor
}

func (m *valueInterceptorFactory[T]) Build(changeValue message.ChangeValue) (res contract.IValueInterceptor, err error) {
	m.once.Do(func() {
		var enumService contract.IEnumService[T]
		if err = m.enumFacotry.Build("ValueType", &enumService); err != nil {
			return
		}

		var valueTypeItems []T
		if valueTypeItems, err = enumService.AllItems(); err != nil {
			return
		}

		for _, r := range valueTypeItems {
			if v, ok := m.valueTypeInterceptor[r.GetValue()]; ok {
				m.metadata[r.GetValue()] = reflect.TypeOf(v)
			} else {
				for _, cr := range m.predicateInterceptors {
					if cr.predicate(r) {
						m.metadata[r.GetValue()] = reflect.TypeOf(cr.interceptor)
						break
					}
				}
			}
		}
	})

	if err != nil {
		return
	}

	if t, ok := m.metadata[changeValue.ValueType]; ok {
		v := reflect.New(t)
		iocsvc.Inject(v, nil)
		res = v.Elem().Interface().(contract.IValueInterceptor)
	} else {
		res = m.defaultInterceptor
	}

	return
}

func (m *valueInterceptorFactory[T]) Register(valueType int, interceptor contract.IValueInterceptor) {
	m.valueTypeInterceptor[valueType] = interceptor
}

func (m *valueInterceptorFactory[T]) RegisterPredicate(predicate func(mcontract.IEnumItem) bool, interceptor contract.IValueInterceptor) {
	m.predicateInterceptors = append(m.predicateInterceptors, predicateInterceptor{
		interceptor: interceptor,
		predicate:   predicate,
	})
}

// 创建枚举拦截器工厂
func NewValueInterceptorFactory[T mcontract.IEnumItem](enumFacotry contract.IEnumFactory) contract.IValueInterceptorFactory {
	return &valueInterceptorFactory[T]{
		defaultInterceptor:    new(emptyValueInterceptor),
		enumFacotry:           enumFacotry,
		metadata:              make(map[int]reflect.Type),
		predicateInterceptors: make([]predicateInterceptor, 0),
		valueTypeInterceptor:  make(map[int]contract.IValueInterceptor),
	}
}
