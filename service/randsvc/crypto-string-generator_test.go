package randsvc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cryptoStringGenerator_Generate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := NewCryptoStringGenerator(6)
		res := self.Generate()
		assert.Len(t, res, 6)
	})
}
