package mongosvc

import (
	"fmt"
	"reflect"
	"sync"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/model/contract"
)

var modelMetadatas sync.Map

type modelMetadata struct {
	idIndex   int
	modelType reflect.Type
	once      sync.Once
	fields    []*fieldMetadata
}

func (m *modelMetadata) FindFields() []*fieldMetadata {
	m.once.Do(func() {
		m.idIndex = -1
		underscore.Range(
			0,
			m.modelType.NumField(),
			1,
		).Map(func(r int, i int) *fieldMetadata {
			f := newFieldMetadata(
				m.modelType.Field(r),
				m.modelType,
			)
			if f.GetTableName() != "" {
				m.idIndex = i
			}
			return f
		}).Value(&(m.fields))
	})

	return m.fields
}

func (m *modelMetadata) GetIDField() (*fieldMetadata, error) {
	fields := m.FindFields()
	if m.idIndex == -1 {
		panic(
			fmt.Sprintf(
				`缺少^alias:"空或自定义表名" db:"列名"^: %s`,
				m.modelType.Name(),
			),
		)
	}

	return fields[m.idIndex], nil
}

func (m *modelMetadata) GetTableName() (string, error) {
	idField, err := m.GetIDField()
	if err != nil {
		return "", err
	}

	return idField.GetTableName(), nil
}

func (m *modelMetadata) GetType() reflect.Type {
	return m.modelType
}

func getModelMetadata[T any](model T) *modelMetadata {
	var t reflect.Type
	if entry, ok := any(model).(contract.IDbModel); ok {
		t = reflect.TypeOf(entry)
	} else if value, ok := any(model).(reflect.Value); ok {
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
		t = value.Type()
	} else {
		t = any(model).(reflect.Type)
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var v any
	var ok bool
	if v, ok = modelMetadatas.Load(t); !ok {
		v = &modelMetadata{
			modelType: t,
		}
		modelMetadatas.Store(t, v)
	}

	return v.(*modelMetadata)
}
