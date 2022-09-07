package contract

import (
	"sync"

	"github.com/ahl5esoft/lite-go/model/contract"
)

type EnumServiceBase[T contract.IEnumItem] struct {
	FindItemsFunc func() ([]T, error)
	mutex         sync.Mutex
	items         []T
}

func (m *EnumServiceBase[T]) AllItems() (res []T, err error) {
	if m.items == nil {
		m.mutex.Lock()
		defer m.mutex.Unlock()

		if m.items == nil {
			if m.items, err = m.FindItemsFunc(); err != nil {
				return
			}
		}
	}

	res = m.items
	return
}

func (m *EnumServiceBase[T]) GetItem(predicate func(T) bool) (res T, err error) {
	var items []T
	if items, err = m.AllItems(); err != nil {
		return
	}

	for _, r := range items {
		if predicate(r) {
			res = r
			break
		}
	}

	return
}
