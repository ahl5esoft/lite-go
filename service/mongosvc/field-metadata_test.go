package mongosvc

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testfieldMetadataModel struct {
	FieldA string `db:"a" alias:""`
	FieldB string `db:"b" alias:"test"`
	FieldC string
}

func (m testfieldMetadataModel) GetID() string {
	return m.FieldA
}

func Test_fieldMetadata_GetColumnName(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		modelType := reflect.TypeOf(testfieldMetadataModel{})
		field, _ := modelType.FieldByName("FieldA")
		self := new(fieldMetadata)
		self.field = field
		self.modelType = modelType
		assert.Equal(
			t,
			self.GetColumnName(),
			"a",
		)
	})

	t.Run("no tag", func(t *testing.T) {
		modelType := reflect.TypeOf(testfieldMetadataModel{})
		field, _ := modelType.FieldByName("FieldC")
		self := new(fieldMetadata)
		self.field = field
		self.modelType = modelType
		assert.Equal(
			t,
			self.GetColumnName(),
			"FieldC",
		)
	})
}

func Test_fieldMetadata_GetTableName(t *testing.T) {
	t.Run("元数据", func(t *testing.T) {
		modelType := reflect.TypeOf(testfieldMetadataModel{})
		field, _ := modelType.FieldByName("FieldB")
		self := new(fieldMetadata)
		self.field = field
		self.modelType = modelType
		assert.Equal(
			t,
			self.GetTableName(),
			"test",
		)
	})

	t.Run("模型名", func(t *testing.T) {
		modelType := reflect.TypeOf(testfieldMetadataModel{})
		field, _ := modelType.FieldByName("FieldA")
		self := new(fieldMetadata)
		self.field = field
		self.modelType = modelType
		assert.Equal(
			t,
			self.GetTableName(),
			modelType.Name(),
		)
	})
}

func Test_fieldMetadata_GetValue(t *testing.T) {
	t.Run("contract.IDbModel", func(t *testing.T) {
		model := testfieldMetadataModel{
			FieldA: "aa",
		}
		modelType := reflect.TypeOf(model)
		field, _ := modelType.FieldByName("FieldA")
		self := new(fieldMetadata)
		self.field = field
		self.modelType = modelType
		assert.Equal(
			t,
			self.GetValue(model),
			model.FieldA,
		)
	})
	t.Run("reflect.Value", func(t *testing.T) {
		model := testfieldMetadataModel{
			FieldA: "aa",
		}
		modelValue := reflect.ValueOf(model)
		field, _ := modelValue.Type().FieldByName("FieldA")
		self := new(fieldMetadata)
		self.field = field
		self.modelType = modelValue.Type()
		assert.Equal(
			t,
			self.GetValue(modelValue),
			model.FieldA,
		)
	})
}
