package contract

import "time"

// 命令接口
type ICommand interface {
	Exec() (stdout string, stderr string, err error)
	SetDir(format string, args ...any) ICommand
	SetExpires(expires time.Duration) ICommand
}
