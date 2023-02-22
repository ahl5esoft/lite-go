package ginsvc

import "github.com/gin-gonic/gin"

type handler struct {
	next         IHandler
	handleAction func(v any) error
}

func (m *handler) SetNext(next IHandler) IHandler {
	m.next = next
	return m
}

func (m *handler) Handle(v any) error {
	if err := m.handleAction(v); err != nil {
		return err
	}

	if m.next == nil {
		return nil
	}

	return m.next.Handle(v)
}

func NewContextHandler(handleAction func(ctx *gin.Context) error) IHandler {
	return &handler{
		handleAction: func(v any) error {
			return handleAction(
				v.(*gin.Context),
			)
		},
	}
}

func NewEngineHandler(handleAction func(app *gin.Engine) error) IHandler {
	return &handler{
		handleAction: func(v any) error {
			return handleAction(
				v.(*gin.Engine),
			)
		},
	}
}
