package iocsvc

import (
	"fmt"
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
)

const (
	instanceIsNotPtr       = "ioc: 注入实例必须是指针"
	invalidTypeFormat      = "ioc: 无效类型(Name = %s, Type = %v)"
	notInterfaceTypeFormat = "ioc: 非接口类型(%v)"
)

var instanceValues = make(map[reflect.Type]map[string]reflect.Value)

// 获取实例
func Get[T any](name string) T {
	return getValueWithName(
		new(T),
		name,
	).Interface().(T)
}

// 是否存在
func Has[T any](name string) bool {
	interfaceType := getInterfaceType(
		new(T),
	)
	if values, ok := instanceValues[interfaceType]; ok {
		if _, ok = values[name]; ok {
			return ok
		}
	}
	return false
}

// 遍历实例内的需要依赖注入的字段进行注入
func Inject(instance interface{}, filterFunc func(reflect.Value) reflect.Value) {
	instanceValue := reflect.ValueOf(instance)
	if instanceValue.Kind() != reflect.Ptr {
		panic(instanceIsNotPtr)
	}

	inject(instanceValue, filterFunc)
}

// 设置依赖注入
func Set[T any](instance T) {
	SetWithName(
		"",
		instance,
	)
}

// 置依赖注入
func SetWithName[T any](name string, instance T) {
	interfaceType := getInterfaceType(
		new(T),
	)
	if _, ok := instanceValues[interfaceType]; !ok {
		instanceValues[interfaceType] = make(map[string]reflect.Value)
	}

	instanceValues[interfaceType][name] = reflect.ValueOf(instance)
}

func getInterfaceType(interfaceObj interface{}) reflect.Type {
	var interfaceType reflect.Type
	var ok bool
	if interfaceType, ok = interfaceObj.(reflect.Type); !ok {
		interfaceType = reflect.TypeOf(interfaceObj)
	}

	if interfaceType.Kind() == reflect.Ptr {
		interfaceType = interfaceType.Elem()
	}

	if interfaceType.Kind() != reflect.Interface {
		panic(
			fmt.Errorf(notInterfaceTypeFormat, interfaceType),
		)
	}

	return interfaceType
}

func getValueWithName(interfaceObj interface{}, name string) reflect.Value {
	interfaceType := getInterfaceType(interfaceObj)
	if values, ok := instanceValues[interfaceType]; ok {
		if v, ok := values[name]; ok {
			return v
		}
	}

	panic(
		fmt.Errorf(invalidTypeFormat, name, interfaceType),
	)
}

func inject(instanceValue reflect.Value, filterFunc func(reflect.Value) reflect.Value) {
	if instanceValue.Kind() == reflect.Ptr {
		instanceValue = instanceValue.Elem()
	}

	underscore.Range(
		0,
		instanceValue.Type().NumField(),
		1,
	).Each(func(r int, _ int) {
		field := instanceValue.Type().Field(r)
		fieldValue := instanceValue.FieldByIndex(field.Index)
		if field.Anonymous {
			if field.Type.Kind() == reflect.Struct {
				inject(fieldValue, filterFunc)
			}
			return
		}

		name, ok := field.Tag.Lookup("inject")
		if !ok {
			return
		}

		if fieldValue.Kind() == reflect.Ptr {
			value := reflect.New(
				field.Type.Elem(),
			)
			fieldValue.Set(value)
			fieldValue = fieldValue.Elem()
		}

		v := getValueWithName(field.Type, name)
		if filterFunc != nil {
			v = filterFunc(v)
		}
		fieldValue.Set(v)
	})
}
