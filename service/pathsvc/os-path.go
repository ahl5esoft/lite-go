package pathsvc

import (
	"path/filepath"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/contract"
)

type osPath struct {
	wd string
}

func (m osPath) Getwd() string {
	return m.wd
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
func NewIOPath(paths ...string) contract.IOsPath {
	p := new(osPath)
	p.wd = p.Join(paths...)
	return p
}
