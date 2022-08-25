package contract

// 相同路径
type IOsPath interface {
	Getwd() string
	Join(paths ...string) string
}
