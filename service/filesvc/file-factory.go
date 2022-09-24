package filesvc

import "github.com/ahl5esoft/lite-go/contract"

type fileFactory struct {
	osPath             contract.IOsPath
	buildFileEntryFunc func(fileFactory contract.IFileFactory, path string) contract.IFileEntry
	buildFileFunc      func(contract.IFileEntry) contract.IFile
}

func (m fileFactory) BuildFile(paths ...string) contract.IFile {
	return m.buildFileFunc(
		m.BuildFileEntry(paths...),
	)
}

func (m fileFactory) BuildFileEntry(paths ...string) contract.IFileEntry {
	return m.buildFileEntryFunc(
		m,
		m.osPath.Join(paths...),
	)
}

// 创建文件工厂
func NewFileFactory(
	osPath contract.IOsPath,
	buildFileEntryFunc func(fileFactory contract.IFileFactory, path string) contract.IFileEntry,
	buildFileFunc func(contract.IFileEntry) contract.IFile,
) contract.IFileFactory {
	return &fileFactory{
		osPath:             osPath,
		buildFileEntryFunc: buildFileEntryFunc,
	}
}
