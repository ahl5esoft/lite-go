package ginsvc

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_NewCallApiContextHandler(t *testing.T) {
	t.Run("resp", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := NewCallApiContextHandler("api", "err", "resp")

		gin.SetMode(gin.DebugMode)
		app := gin.New()
		app.POST("/", func(ctx *gin.Context) {
			mockApi := contract.NewMockIApi(ctrl)
			ctx.Set("api", mockApi)

			mockApi.EXPECT().Call().Return("ok", nil)

			err := self.Handle(ctx)
			assert.Nil(t, err)

			res := ctx.GetString("resp")
			assert.Equal(t, res, "ok")
		})

		req := httptest.NewRequest(
			"POST",
			"/",
			strings.NewReader(""),
		)
		resp := httptest.NewRecorder()
		app.ServeHTTP(resp, req)
	})

	t.Run("err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := NewCallApiContextHandler("api", "err", "resp")

		gin.SetMode(gin.DebugMode)
		app := gin.New()
		app.POST("/", func(ctx *gin.Context) {
			mockApi := contract.NewMockIApi(ctrl)
			ctx.Set("api", mockApi)

			cErr := errors.New("err")
			mockApi.EXPECT().Call().Return("ok", cErr)

			err := self.Handle(ctx)
			assert.Nil(t, err)

			res := ctx.GetString("resp")
			assert.Empty(t, res)
			v, ok := ctx.Get("err")
			assert.True(t, ok)
			assert.Equal(t, v, cErr)
		})

		req := httptest.NewRequest(
			"POST",
			"/",
			strings.NewReader(""),
		)
		resp := httptest.NewRecorder()
		app.ServeHTTP(resp, req)
	})
}
