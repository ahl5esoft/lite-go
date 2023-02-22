package ginsvc

import (
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testGetApi struct{}

func (m testGetApi) Call() (any, error) {
	return nil, nil
}

func Test_NewGetApiContextHandler(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockError := contract.NewMockIError(ctrl)
		self := NewGetApiContextHandler(mockError, "api", "api-name", "end", "err")

		gin.SetMode(gin.DebugMode)
		app := gin.New()
		app.POST("/", func(ctx *gin.Context) {
			ctx.Set("api-name", "bb")
			ctx.Set("end", "aa")

			err := self.Handle(ctx)
			assert.Nil(t, err)

			_, ok := ctx.Get("api")
			assert.False(t, ok)

			v, ok := ctx.Get("err")
			assert.True(t, ok)
			assert.Equal(t, v, mockError)
		})

		req := httptest.NewRequest(
			"POST",
			"/",
			strings.NewReader(""),
		)
		resp := httptest.NewRecorder()
		app.ServeHTTP(resp, req)
	})

	t.Run("ok", func(t *testing.T) {
		self := NewGetApiContextHandler(nil, "api", "api-name", "end", "err")

		RegisterApi("aa", "bb", testGetApi{})

		gin.SetMode(gin.DebugMode)
		app := gin.New()
		app.POST("/", func(ctx *gin.Context) {
			ctx.Set("api-name", "bb")
			ctx.Set("end", "aa")

			err := self.Handle(ctx)
			assert.Nil(t, err)

			api, ok := ctx.Get("api")
			assert.True(t, ok)
			apiType := reflect.TypeOf(api).Elem()
			assert.Equal(
				t,
				apiType,
				reflect.TypeOf(testGetApi{}),
			)
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
