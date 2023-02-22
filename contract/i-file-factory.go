package contract

type IFileFactory interface {
	BuildFile(paths ...string) IFile
	BuildFileEntry(paths ...string) IFileEntry
}
