package cmdsvc

import "github.com/ahl5esoft/lite-go/contract"

type commandFactory func(string, []string) contract.ICommand

func (m commandFactory) Build(name string, args ...string) contract.ICommand {
	return m(name, args)
}

// 创建命令工厂
func NewCommandFactory(buildFunc func(string, []string) contract.ICommand) contract.ICommandFactory {
	return commandFactory(buildFunc)
}
