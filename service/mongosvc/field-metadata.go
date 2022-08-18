package mongosvc

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/model/contract"
)

type fieldMetadata struct {
	field      reflect.StructField
	modelType  reflect.Type
	columnName string
	tableName  *string
}

func (m *fieldMetadata) GetColumnName() string {
	if m.columnName == "" {
		m.columnName = m.field.Tag.Get("db")
		if m.columnName == "" {
			m.columnName = m.field.Name
		}
	}

	return m.columnName
}

func (m *fieldMetadata) GetTableName() string {
	if m.tableName == nil {
		v, ok := m.field.Tag.Lookup("alias")
		if ok && v == "" {
			v = m.modelType.Name()
		}

		m.tableName = &v
	}

	return *m.tableName
}

func (m *fieldMetadata) GetValue(v any) interface{} {
	var value reflect.Value
	if entry, ok := v.(contract.IDbModel); ok {
		value = reflect.ValueOf(entry)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
	} else {
		value = v.(reflect.Value)
	}
	return value.FieldByIndex(m.field.Index).Interface()
}

func newFieldMetadata(field reflect.StructField, modelType reflect.Type) *fieldMetadata {
	return &fieldMetadata{
		field:     field,
		modelType: modelType,
	}
}
