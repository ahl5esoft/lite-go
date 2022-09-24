package contract

// 文件工厂
type IFileFactory interface {
	// 创建文件
	BuildFile(paths ...string) IFile
	// 创建文件项(文件或目录)
	BuildFileEntry(paths ...string) IFileEntry
}
