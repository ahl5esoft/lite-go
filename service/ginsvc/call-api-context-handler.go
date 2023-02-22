package ginsvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	"github.com/gin-gonic/gin"
)

func NewCallApiContextHandler(ctxApiKey, ctxErrKey, ctxRespKey string) IHandler {
	return NewContextHandler(func(ctx *gin.Context) error {
		if v, ok := ctx.Get(ctxApiKey); ok {
			if resp, err := v.(contract.IApi).Call(); err != nil {
				ctx.Set(ctxErrKey, err)
			} else {
				ctx.Set(ctxRespKey, resp)
			}
		}
		return nil
	})
}
