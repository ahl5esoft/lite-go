package ginsvc

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/gin-gonic/gin"
)

var (
	apiMetadata = make(map[string]map[string]reflect.Type)
)

func NewGetApiContextHandler(apiNilErr contract.IError, ctxApiKey, ctxApiNameKey, ctxEndpointKey, ctxErrKey string) IHandler {
	return NewContextHandler(func(ctx *gin.Context) error {
		var api contract.IApi
		endpoint := ctx.GetString(ctxEndpointKey)
		if apiTypes, ok := apiMetadata[endpoint]; ok {
			apiName := ctx.GetString(ctxApiNameKey)
			if apiType, ok := apiTypes[apiName]; ok {
				api = reflect.New(apiType).Interface().(contract.IApi)
			}
		}

		if api == nil {
			ctx.Set(ctxErrKey, apiNilErr)
		} else {
			ctx.Set(ctxApiKey, api)
		}

		return nil
	})
}

func RegisterApi(endpoint, name string, api contract.IApi) {
	if _, ok := apiMetadata[endpoint]; !ok {
		apiMetadata[endpoint] = make(map[string]reflect.Type)
	}

	apiType := reflect.TypeOf(api)
	if apiType.Kind() == reflect.Ptr {
		apiType = apiType.Elem()
	}
	apiMetadata[endpoint][name] = apiType
}
