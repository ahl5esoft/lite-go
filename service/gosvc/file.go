package gosvc

import (
	"fmt"
	"os"
	"reflect"

	"github.com/ahl5esoft/lite-go/contract"
)

type file struct {
	contract.IFileEntry
}

func (m file) Read(v any) (err error) {
	var bf []byte
	bf, err = os.ReadFile(
		m.GetPath(),
	)
	if err != nil {
		return
	}

	rv := reflect.ValueOf(v).Elem()
	if _, ok := v.(*string); ok {
		rv.SetString(
			string(bf),
		)
	} else if _, ok := v.(*[]byte); ok {
		rv.SetBytes(bf)
	} else {
		err = fmt.Errorf(
			"暂不支持IFile.Read(%s)",
			rv.Type(),
		)
	}

	return
}

func newFile(entry contract.IFileEntry) contract.IFile {
	return &file{
		IFileEntry: entry,
	}
}
