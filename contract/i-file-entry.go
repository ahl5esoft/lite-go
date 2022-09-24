package contract

// 文件项
type IFileEntry interface {
	// 拷贝到
	CopyTo(...string) error
	// 是否存在
	Exists() bool
	// 获取路径
	GetPath() string
	// 移动到
	MoveTo(...string) error
	// 删除
	Remove() error
}
