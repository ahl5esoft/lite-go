package ginsvc

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_apiPort_Listen(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	self := new(apiPort)

	req := httptest.NewRequest(
		"POST",
		"/endpoint/api",
		strings.NewReader(""),
	)
	resp := httptest.NewRecorder()
	self.options = []Option{
		func(app *gin.Engine) {
			app.POST("/endpoint/api", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, map[string]interface{}{
					"data": "ok",
					"err":  0,
				})
			})
		},
		func(app *gin.Engine) {
			app.ServeHTTP(resp, req)
		},
	}

	self.Listen()

	res := resp.Result()
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	assert.Equal(
		t,
		string(body),
		`{"data":"ok","err":0}`,
	)
}
