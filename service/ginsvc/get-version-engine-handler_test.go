package ginsvc

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_NewGetVersionEngineHandler(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockFileFactory := contract.NewMockIFileFactory(ctrl)
		mockFilePath := contract.NewMockIFilePath(ctrl)
		self := NewGetVersionEngineHandler(mockFileFactory, mockFilePath, "test")

		mockFilePath.EXPECT().Getwd().Return("wd")

		mockFile := contract.NewMockIFile(ctrl)
		mockFileFactory.EXPECT().BuildFile("wd", "README.md").Return(mockFile)

		mockFile.EXPECT().Read(
			gomock.Any(),
		).SetArg(0, "v1.2.3")

		mockHandler := NewMockIHandler(ctrl)
		self.SetNext(mockHandler)

		gin.SetMode(gin.DebugMode)
		app := gin.New()
		mockHandler.EXPECT().Handle(app)

		err := self.Handle(app)
		assert.Nil(t, err)

		req := httptest.NewRequest(
			http.MethodGet,
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
			`{"data":{"project":"test","version":"1.2.3"},"err":0}`,
		)
	})
}
