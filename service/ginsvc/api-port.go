package ginsvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	"github.com/gin-gonic/gin"
)

type apiPort struct {
	handler IHandler
	mode    string
}

func (m apiPort) Listen() {
	gin.SetMode(m.mode)
	app := gin.New()
	if err := m.handler.Handle(app); err != nil {
		panic(err)
	}
}

func NewApiPort(handler IHandler, mode string) contract.IApiPort {
	return &apiPort{
		handler: handler,
		mode:    mode,
	}
}
