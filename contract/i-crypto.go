package contract

// 密码
type ICrypto interface {
	// 比较
	Compare([]byte, []byte) (bool, error)
	// 解密
	Decrypt([]byte) ([]byte, error)
	// 加密
	Encrypt([]byte) ([]byte, error)
}
