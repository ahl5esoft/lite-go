package randsvc

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/ahl5esoft/lite-go/contract"
)

type cryptoStringGenerator int

func (m cryptoStringGenerator) Generate() string {
	b := make([]byte, m)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:m]
}

func NewCryptoStringGenerator(len int) contract.IStringGenerator {
	return cryptoStringGenerator(len)
}
