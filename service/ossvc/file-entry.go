package ossvc

import (
	"os"

	"github.com/ahl5esoft/lite-go/contract"
)

type fileEntry struct {
	fileFactory contract.IFileFactory
	osPath      contract.IOsPath
	path        string
}

func (m fileEntry) CopyTo(paths ...string) error {
	dst := m.fileFactory.BuildFileEntry(paths...)
	if dst.Exists() {
		return os.ErrExist
	}

	return m.MoveTo(
		dst.GetPath(),
	)
}

func (m fileEntry) Exists() bool {
	_, err := os.Stat(m.path)
	return err == nil || os.IsExist(err)
}

func (m fileEntry) GetPath() string {
	return m.path
}

func (m fileEntry) MoveTo(paths ...string) error {
	dstPath := m.osPath.Join(paths...)
	return os.Rename(m.path, dstPath)
}

func (m fileEntry) Remove() error {
	return os.RemoveAll(m.path)
}

// 创建文件项
func NewFileEntry(
	fileFactory contract.IFileFactory,
	osPath contract.IOsPath,
	path string,
) contract.IFileEntry {
	return fileEntry{
		fileFactory: fileFactory,
		osPath:      osPath,
		path:        path,
	}
}
