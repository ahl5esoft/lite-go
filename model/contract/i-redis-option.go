package contract

// redis选项
type IRedisOption interface {
	// 地址
	GetAddr() string
	// 密码
	GetPassword() string
}
