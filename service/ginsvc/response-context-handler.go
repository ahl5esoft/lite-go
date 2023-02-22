package ginsvc

import (
	"fmt"
	"net/http"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/model/message"
	"github.com/gin-gonic/gin"
)

func NewResponseContextHandler(ctxErrKey, ctxRespKey string) IHandler {
	return NewContextHandler(func(ctx *gin.Context) error {
		var err error
		if v, ok := ctx.Get(ctxErrKey); ok {
			var ok bool
			if err, ok = v.(error); !ok {
				err = fmt.Errorf("panic: %+v", v)
			}
		}

		var resp message.ApiResponse
		if err != nil {
			if cErr, ok := err.(contract.IError); ok {
				resp.Data = cErr.GetData()
				resp.Error = cErr.GetCode()
			} else {
				fmt.Println(err)
				resp.Error = 599
			}
		} else {
			resp.Data = ctx.MustGet(ctxRespKey)
		}
		ctx.JSON(http.StatusOK, resp)
		return nil
	})
}
