package ginsvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"
	headerkey "github.com/ahl5esoft/lite-go/model/enum/header-key"
	"github.com/ahl5esoft/lite-go/model/message"
	"github.com/ahl5esoft/lite-go/service/errorsvc"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

type ApiSession struct {
	Crypto   contract.ICrypto `inject:"auth-crypto"`
	UserAuth message.UserAuth
}

func (m *ApiSession) SetSession(ctx *gin.Context) error {
	if s := ctx.GetHeader(headerkey.AuthData); s != "" {
		bf, err := m.Crypto.Decrypt(
			[]byte(s),
		)
		if err != nil {
			return err
		}

		if err = jsoniter.Unmarshal(bf, &(m.UserAuth)); err == nil {
			return nil
		}
	}

	return errorsvc.Newf(errorcode.Auth, "")
}
