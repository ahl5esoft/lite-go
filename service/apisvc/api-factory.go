package apisvc

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/contract"
	errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"
	"github.com/ahl5esoft/lite-go/service/errorsvc"
)

var errNilApi = errorsvc.New(errorcode.API, nil)

// api工厂
type apiFactory map[string]map[string]reflect.Type

// 创建api
func (m apiFactory) Build(endpoint, api string) (instance contract.IApi, err error) {
	if apiTypes, ok := m[endpoint]; ok {
		if apiType, ok := apiTypes[api]; ok {
			instance = reflect.New(apiType).Interface().(contract.IApi)
			return
		}
	}

	err = errNilApi
	return
}

// 注册
func (m apiFactory) Register(endpoint, name string, api any) {
	if _, ok := m[endpoint]; !ok {
		m[endpoint] = make(map[string]reflect.Type)
	}

	apiType := reflect.TypeOf(api)
	if apiType.Kind() == reflect.Ptr {
		apiType = apiType.Elem()
	}
	m[endpoint][name] = apiType
}

// 创建api工厂
func NewApiFactory() contract.IApiFactory {
	return make(apiFactory)
}
