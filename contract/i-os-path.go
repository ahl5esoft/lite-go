package contract

// 相同路径
type IOsPath interface {
	GetRoot() string
	Join(paths ...string) string
}
