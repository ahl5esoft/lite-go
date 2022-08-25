package ossvc

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fileEntry_Exists(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		wd, _ := os.Getwd()
		res := NewFileEntry(
			nil,
			nil,
			path.Join(wd, "file.go"),
		).Exists()
		assert.True(t, res)
	})
}
