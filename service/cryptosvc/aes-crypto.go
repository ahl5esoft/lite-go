package cryptosvc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"

	"github.com/ahl5esoft/lite-go/contract"
)

type aesCrypto []byte

func (m aesCrypto) Compare(_ []byte, _ []byte) (bool, error) {
	return false, nil
}

func (m aesCrypto) Decrypt(cipherText []byte) ([]byte, error) {
	return m.getCtrCrypt(cipherText)
}

func (m aesCrypto) Encrypt(plaintext []byte) ([]byte, error) {
	return m.getCtrCrypt(plaintext)
}

func (m aesCrypto) getCtrCrypt(text []byte) (res []byte, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher(m); err != nil {
		return
	}

	iv := bytes.Repeat(
		[]byte("1"),
		block.BlockSize(),
	)
	res = make(
		[]byte,
		len(text),
	)
	cipher.NewCTR(block, iv).XORKeyStream(res, text)
	return

}

// 创建aes
func NewAesCrypto(key []byte) contract.ICrypto {
	return aesCrypto(key)
}
