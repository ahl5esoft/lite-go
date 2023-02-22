package gosvc

import (
	fp "path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_filePath_Join(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		res := new(filePath).Join("a", "b")
		assert.Equal(
			t,
			res,
			fp.Join("a", "b"),
		)
	})

	t.Run("HasParent", func(t *testing.T) {
		res := new(filePath).Join("a", "b", "..", "c")
		assert.Equal(
			t,
			res,
			fp.Join("a", "c"),
		)
	})

	t.Run("/", func(t *testing.T) {
		res := new(filePath).Join("a", "b/c/d")
		assert.Equal(
			t,
			res,
			fp.Join("a", "b", "c", "d"),
		)
	})
}
