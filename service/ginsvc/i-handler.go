package ginsvc

type IHandler interface {
	Handle(v any) error
	SetNext(next IHandler) IHandler
}
