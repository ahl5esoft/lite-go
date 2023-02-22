package ginsvc

import (
	"github.com/gin-gonic/gin"
)

func NewPostEngineHandler(apiHandler, responseHandler IHandler, ctxErrKey, rule string) IHandler {
	return NewEngineHandler(func(app *gin.Engine) error {
		app.POST(rule, func(ctx *gin.Context) {
			defer func() {
				if rv := recover(); rv != nil {
					ctx.Set(ctxErrKey, rv)
				}

				responseHandler.Handle(ctx)
			}()

			if err := apiHandler.Handle(ctx); err != nil {
				ctx.Set(ctxErrKey, err)
			}
		})
		return nil
	})
}
