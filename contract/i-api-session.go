package contract

// api会话
type IApiSession[T any] interface {
	SetSession(T) error
}
