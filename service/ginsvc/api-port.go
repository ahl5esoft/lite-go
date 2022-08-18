package ginsvc

import (
	"github.com/ahl5esoft/lite-go/contract"

	"github.com/gin-gonic/gin"
)

type apiPort struct {
	options []Option
}

func (m apiPort) Listen() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	for _, r := range m.options {
		r(app)
	}
}

func NewApiPort(options ...Option) contract.IApiPort {
	return &apiPort{
		options: options,
	}
}
