package ginsvc

import (
	"net/url"

	errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"
	headerkey "github.com/ahl5esoft/lite-go/model/enum/header-key"
	"github.com/ahl5esoft/lite-go/model/message"
	"github.com/ahl5esoft/lite-go/service/errorsvc"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

type ApiSession struct {
	UserAuth message.UserAuth
}

func (m *ApiSession) SetSession(ctx *gin.Context) error {
	v := ctx.GetHeader(headerkey.AuthData)
	if v != "" {
		var err error
		if v, err = url.QueryUnescape(v); err == nil {
			if err = jsoniter.UnmarshalFromString(v, &(m.UserAuth)); err == nil {
				return nil
			}
		}
	}

	return errorsvc.Newf(errorcode.Auth, "")
}
