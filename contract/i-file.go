package contract

// 文件
type IFile interface {
	// 读取
	Read(any) error
}
