package gosvc

import "github.com/ahl5esoft/lite-go/contract"

type fileFactory struct {
	filePath contract.IFilePath
}

func (m fileFactory) BuildFile(paths ...string) contract.IFile {
	fileEntry := m.BuildFileEntry(paths...)
	return newFile(fileEntry)
}

func (m fileFactory) BuildFileEntry(paths ...string) contract.IFileEntry {
	return newFileEntry(
		m,
		m.filePath,
		m.filePath.Join(paths...),
	)
}

func NewFileFactory(
	filePath contract.IFilePath,
) contract.IFileFactory {
	return &fileFactory{
		filePath: filePath,
	}
}
