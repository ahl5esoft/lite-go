package contract

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/model/contract"
	"github.com/ahl5esoft/lite-go/model/message"
)

type ValueServiceBase[T contract.IValue] struct {
	FindEntriesFunc func() ([]T, error)
	UpdateFunc      func(IUnitOfWork, string, []message.ChangeValue, *[]T) error
	entries         []T
}

func (m *ValueServiceBase[T]) CheckConditions(_ IUnitOfWork, _ [][]contract.IValueCondition) (bool, error) {
	return false, nil
}

func (m *ValueServiceBase[T]) GetCount(uow IUnitOfWork, valueType int) (res int64, err error) {
	var entries []T
	if entries, err = m.findEntries(); err != nil {
		return
	}

	if len(entries) > 0 {
		if v, ok := entries[0].GetValue()[valueType]; ok {
			res = v
		}
	}

	return
}

func (m *ValueServiceBase[T]) GetEntry(v any) (err error) {
	var entries []T
	if entries, err = m.findEntries(); err != nil {
		return
	} else if len(entries) > 0 {
		reflect.ValueOf(v).Elem().Set(
			reflect.ValueOf(entries[0]),
		)
	}

	return
}

func (m *ValueServiceBase[T]) Update(uow IUnitOfWork, source string, changeValues []message.ChangeValue) (err error) {
	if _, err = m.findEntries(); err != nil {
		return
	}

	err = m.UpdateFunc(uow, source, changeValues, &(m.entries))
	return
}

func (m *ValueServiceBase[T]) findEntries() (res []T, err error) {
	if m.entries == nil {
		if m.entries, err = m.FindEntriesFunc(); err != nil {
			return
		}
	}

	res = m.entries
	return
}
