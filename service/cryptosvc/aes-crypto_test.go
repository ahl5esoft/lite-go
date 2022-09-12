package cryptosvc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_aesCrypto_Decrypt(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		res, err := NewAesCrypto(
			[]byte("a123b456c789d123"),
		).Decrypt(
			[]byte{0x24, 0x86, 0xf7, 0x59, 0x81},
		)
		assert.Nil(t, err)
		assert.EqualValues(t, res, []byte("hello"))
	})
}

func Test_aesCrypto_Encrypt(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		res, err := NewAesCrypto(
			[]byte("a123b456c789d123"),
		).Encrypt(
			[]byte("hello"),
		)
		assert.Nil(t, err)
		assert.EqualValues(t, res, []byte{0x24, 0x86, 0xf7, 0x59, 0x81})
	})
}
