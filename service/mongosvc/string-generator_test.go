package mongosvc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringGenerator_Generate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := NewStringGenerator()
		res := self.Generate()
		assert.Len(t, res, 24)
	})
}
