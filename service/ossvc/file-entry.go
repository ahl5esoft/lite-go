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

func (m fileEntry) Copy(paths ...string) error {
	fileEntry := m.fileFactory.BuildFileEntry(paths...)
	if err := fileEntry.Remove(); err != nil {
		return err
	}

	return m.MoveTo(
		fileEntry.GetPath(),
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
	if !m.Exists() {
		return nil
	}

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
