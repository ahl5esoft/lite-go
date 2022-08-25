package contract

// 命令工厂
type ICommandFactory interface {
	// 创建命令
	Build(name string, args ...string) ICommand
}
