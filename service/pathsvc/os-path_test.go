package pathsvc

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_osPath_Join(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		res := new(osPath).Join("a", "b")
		assert.Equal(
			t,
			res,
			filepath.Join("a", "b"),
		)
	})

	t.Run("HasParent", func(t *testing.T) {
		res := new(osPath).Join("a", "b", "..", "c")
		assert.Equal(
			t,
			res,
			filepath.Join("a", "c"),
		)
	})

	t.Run("/", func(t *testing.T) {
		res := new(osPath).Join("a", "b/c/d")
		assert.Equal(
			t,
			res,
			filepath.Join("a", "b", "c", "d"),
		)
	})
}
