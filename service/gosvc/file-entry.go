package gosvc

import (
	"os"

	"github.com/ahl5esoft/lite-go/contract"
)

type fileEntry struct {
	fileFactory contract.IFileFactory
	filePath    contract.IFilePath
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
	dstPath := m.filePath.Join(paths...)
	return os.Rename(m.path, dstPath)
}

func (m fileEntry) Remove() error {
	return os.RemoveAll(m.path)
}

func newFileEntry(
	fileFactory contract.IFileFactory,
	filePath contract.IFilePath,
	path string,
) contract.IFileEntry {
	return fileEntry{
		fileFactory: fileFactory,
		filePath:    filePath,
		path:        path,
	}
}
