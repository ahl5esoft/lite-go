package execsvc

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/ahl5esoft/lite-go/contract"
)

var errExecTimeout = fmt.Errorf("执行超时")

type command struct {
	name, wd string
	expires  time.Duration
	args     []string
}

func (m *command) Exec() (string, string, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		m.expires,
	)
	defer func() {
		m.expires = time.Second * 5
		m.wd = ""
		cancel()
	}()

	cmd := exec.CommandContext(ctx, m.name, m.args...)
	if m.wd != "" {
		cmd.Dir = m.wd
	}

	var stderrBf bytes.Buffer
	cmd.Stderr = &stderrBf

	var stdoutBf bytes.Buffer
	cmd.Stdout = &stdoutBf

	if err := cmd.Start(); err != nil {
		return stdoutBf.String(), stderrBf.String(), err
	}

	err := cmd.Wait()
	if ctx.Err() == context.DeadlineExceeded {
		err = errExecTimeout
	}
	return stdoutBf.String(), stderrBf.String(), err
}

func (m *command) SetDir(format string, args ...interface{}) contract.ICommand {
	m.wd = fmt.Sprintf(format, args...)
	return m
}

func (m *command) SetExpires(expires time.Duration) contract.ICommand {
	m.expires = expires
	return m
}

// 创建ICommand
func NewCommand(name string, args []string) contract.ICommand {
	return &command{
		args:    args,
		name:    name,
		expires: time.Second * 5,
	}
}
