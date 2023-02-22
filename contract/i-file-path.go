package contract

type IFilePath interface {
	Getwd() string
	Join(paths ...string) string
}
