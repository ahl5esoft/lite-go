package randsvc

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/ahl5esoft/lite-go/contract"
)

type cryptoStringGenerator int

func (m cryptoStringGenerator) Generate() string {
	b := make([]byte, m)
	rand.Read(b)
	fmt.Println(len(b))
	return base64.URLEncoding.EncodeToString(b)
}

func NewCryptoStringGenerator(len int) contract.IStringGenerator {
	return cryptoStringGenerator(len)
}
