package iocsvc

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ahl5esoft/lite-go/service/reflectsvc"

	"github.com/stretchr/testify/assert"
)

type iInterface interface {
	Test()
}

type derive struct{}

func (m derive) Test() {
	fmt.Println("set test")
}

type defaultTest struct {
	One iInterface `inject:""`
}

type customTest struct {
	One iInterface `inject:"custom"`
}

type composeTest struct {
	defaultTest

	Child iInterface `inject:"custom"`
}

type composeInterfaceTest struct {
	iInterface

	Child iInterface `inject:"custom"`
}

func Test_Get(t *testing.T) {
	t.Run("无效", func(t *testing.T) {
		ct := getInterfaceType(
			new(iInterface),
		)
		defer func() {
			rv := recover()
			assert.NotNil(t, rv)

			err, ok := rv.(error)
			assert.True(t, ok)
			assert.Equal(
				t,
				err,
				fmt.Errorf(invalidTypeFormat, "", ct),
			)
		}()

		Get[iInterface]("")
	})
}

func Test_Has(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ct := getInterfaceType(
			new(iInterface),
		)
		instanceValues[ct] = map[string]reflect.Value{
			"": reflect.ValueOf(1),
		}
		defer delete(instanceValues, ct)

		assert.True(
			t,
			Has[iInterface](""),
		)
	})

	t.Run("false", func(t *testing.T) {
		assert.False(
			t,
			Has[iInterface](""),
		)
	})
}

func Test_Inject(t *testing.T) {
	t.Run("默认", func(t *testing.T) {
		it := getInterfaceType(
			new(iInterface),
		)
		instanceValues[it] = map[string]reflect.Value{
			"": reflect.ValueOf(
				new(derive),
			),
		}

		var m defaultTest
		Inject(&m, func(v reflect.Value) reflect.Value {
			return v
		})

		assert.Equal(
			t,
			m.One,
			instanceValues[it][""].Interface(),
		)
	})

	t.Run("自定义", func(t *testing.T) {
		it := getInterfaceType(
			new(iInterface),
		)
		instanceValues[it] = map[string]reflect.Value{
			"custom": reflect.ValueOf(
				new(derive),
			),
		}

		var m customTest
		Inject(&m, func(v reflect.Value) reflect.Value {
			return v
		})

		assert.Equal(
			t,
			m.One,
			instanceValues[it]["custom"].Interface(),
		)
	})

	t.Run("filterFunc is nil", func(t *testing.T) {
		it := getInterfaceType(
			new(iInterface),
		)
		instanceValues[it] = map[string]reflect.Value{
			"": reflect.ValueOf(
				new(derive),
			),
		}

		var m defaultTest
		Inject(&m, nil)

		assert.Equal(
			t,
			m.One,
			instanceValues[it][""].Interface(),
		)
	})

	t.Run("继承", func(t *testing.T) {
		it := getInterfaceType(
			new(iInterface),
		)
		instanceValues[it] = map[string]reflect.Value{
			"": reflect.ValueOf(
				new(derive),
			),
			"custom": reflect.ValueOf(
				new(derive),
			),
		}

		var self composeTest
		Inject(&self, nil)

		assert.Equal(
			t,
			self.Child,
			instanceValues[it][""].Interface(),
		)
		assert.Equal(
			t,
			self.One,
			instanceValues[it][""].Interface(),
		)
	})

	t.Run("继承接口", func(t *testing.T) {
		it := getInterfaceType(
			new(iInterface),
		)
		instanceValues[it] = map[string]reflect.Value{
			"": reflect.ValueOf(
				new(derive),
			),
			"custom": reflect.ValueOf(
				new(derive),
			),
		}

		var self composeInterfaceTest
		Inject(&self, nil)

		assert.Equal(
			t,
			self.Child,
			instanceValues[it][""].Interface(),
		)
	})
}

func Test_Set(t *testing.T) {
	t.Run("默认", func(t *testing.T) {
		ct := reflectsvc.InterfaceTypeOf(
			new(iInterface),
		)
		defer delete(instanceValues, ct)

		Set[iInterface](
			new(derive),
		)

		values, ok := instanceValues[ct]
		assert.True(t, ok)

		_, ok = values[""]
		assert.True(t, ok)
	})
}

func Test_SetWithName(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ct := reflectsvc.InterfaceTypeOf(
			new(iInterface),
		)
		defer delete(instanceValues, ct)

		SetWithName[iInterface](
			"c",
			new(derive),
		)

		values, ok := instanceValues[ct]
		assert.True(t, ok)

		_, ok = values["c"]
		assert.True(t, ok)
	})
}
