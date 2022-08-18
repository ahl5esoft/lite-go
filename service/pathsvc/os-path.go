package pathsvc

import (
	"path/filepath"

	"github.com/ahl5esoft/lite-go/contract"

	underscore "github.com/ahl5esoft/golang-underscore"
)

type osPath struct {
	root string
}

func (m osPath) GetRoot() string {
	return m.root
}

func (m osPath) Join(paths ...string) string {
	var res string
	underscore.Chain(paths).Aggregate(func(memo string, r string, _ int) string {
		if memo == "" {
			return r
		}

		if r == ".." {
			return filepath.Dir(memo)
		}

		return filepath.Join(memo, r)
	}, "").Value(&res)
	return res
}

// 创建系统路径
func NewIOPath(rootArgs ...string) contract.IOsPath {
	p := new(osPath)
	p.root = p.Join(rootArgs...)
	return p
}
