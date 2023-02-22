package ginsvc

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func Test_NewPostEngineHandler(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockApiHandler := NewMockIHandler(ctrl)
		mockRespHandler := NewMockIHandler(ctrl)
		self := NewPostEngineHandler(mockApiHandler, mockRespHandler, "err", "/")

		mockApiHandler.EXPECT().Handle(
			gomock.Any(),
		)

		mockRespHandler.EXPECT().Handle(
			gomock.Any(),
		)

		gin.SetMode(gin.DebugMode)
		app := gin.New()
		self.Handle(app)

		req := httptest.NewRequest(
			"POST",
			"/",
			strings.NewReader(""),
		)
		resp := httptest.NewRecorder()
		app.ServeHTTP(resp, req)
	})
}
