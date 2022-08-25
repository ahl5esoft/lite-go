package filesvc

import "github.com/ahl5esoft/lite-go/contract"

type fileFactory struct {
	osPath             contract.IOsPath
	buildFileEntryFunc func(contract.IFileFactory, contract.IOsPath, string) contract.IFileEntry
}

func (m fileFactory) BuildFileEntry(paths ...string) contract.IFileEntry {
	return m.buildFileEntryFunc(
		m,
		m.osPath,
		m.osPath.Join(paths...),
	)
}

// 创建文件工厂
func NewFileFactory(
	osPath contract.IOsPath,
	buildFileEntryFunc func(contract.IFileFactory, contract.IOsPath, string) contract.IFileEntry,
) contract.IFileFactory {
	return &fileFactory{
		osPath:             osPath,
		buildFileEntryFunc: buildFileEntryFunc,
	}
}
