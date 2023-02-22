package ginsvc

import (
	"errors"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type testError struct {
	code errorcode.Value
	data any
}

func (m testError) GetData() any {
	return m.data
}

func (m testError) GetCode() errorcode.Value {
	return m.code
}

func (m testError) Error() string {
	return ""
}

func Test_NewResponseContextHandler(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		self := NewResponseContextHandler("err", "resp")

		gin.SetMode(gin.DebugMode)
		app := gin.New()
		app.POST("/", func(ctx *gin.Context) {
			defer self.Handle(ctx)

			ctx.Set(
				"err",
				errors.New("500"),
			)
		})

		req := httptest.NewRequest(
			"POST",
			"/",
			strings.NewReader(""),
		)

		resp := httptest.NewRecorder()
		app.ServeHTTP(resp, req)

		res := resp.Result()
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)
		assert.Equal(
			t,
			string(body),
			`{"data":null,"err":599}`,
		)
	})

	t.Run("IError", func(t *testing.T) {
		self := NewResponseContextHandler("err", "resp")

		gin.SetMode(gin.DebugMode)
		app := gin.New()
		app.POST("/", func(ctx *gin.Context) {
			defer self.Handle(ctx)

			ctx.Set("err", testError{
				code: 11,
				data: "tt",
			})
		})

		req := httptest.NewRequest(
			"POST",
			"/",
			strings.NewReader(""),
		)

		resp := httptest.NewRecorder()
		app.ServeHTTP(resp, req)

		res := resp.Result()
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)
		assert.Equal(
			t,
			string(body),
			`{"data":"tt","err":11}`,
		)
	})

	t.Run("resp", func(t *testing.T) {
		self := NewResponseContextHandler("err", "resp")

		gin.SetMode(gin.DebugMode)
		app := gin.New()
		app.POST("/", func(ctx *gin.Context) {
			defer self.Handle(ctx)

			ctx.Set("resp", "rr")
		})

		req := httptest.NewRequest(
			"POST",
			"/",
			strings.NewReader(""),
		)

		resp := httptest.NewRecorder()
		app.ServeHTTP(resp, req)

		res := resp.Result()
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)
		assert.Equal(
			t,
			string(body),
			`{"data":"rr","err":0}`,
		)
	})
}
