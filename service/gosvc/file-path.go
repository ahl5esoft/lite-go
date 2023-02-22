package gosvc

import (
	"path/filepath"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/contract"
)

type filePath struct {
	wd string
}

func (m filePath) Getwd() string {
	return m.wd
}

func (m filePath) Join(paths ...string) string {
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

func NewFilePath(paths ...string) contract.IFilePath {
	p := new(filePath)
	p.wd = p.Join(paths...)
	return p
}
