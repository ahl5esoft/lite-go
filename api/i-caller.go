//go:generate mockgen -destination i-caller_mock.go -package api github.com/ahl5esoft/lite-go/api ICaller

package api

import "time"

// ICaller is api调用接口
type ICaller interface {
	Call(route string, body interface{}, expires time.Duration) (interface{}, error)
	VoidCall(route string, body interface{}) error
}
