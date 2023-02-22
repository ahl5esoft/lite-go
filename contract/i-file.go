package contract

// 文件
type IFile interface {
	IFileEntry

	// 读取
	Read(any) error
}
