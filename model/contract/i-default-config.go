package contract

// 默认配置
type IDefaultConfig interface {
	// 认证密钥
	GetAuthCipher() string
	// 网关
	GetGateway() string
	// mongo
	GetMongo() string
	// 项目名
	GetName() string
	// 端口
	GetPort() int
	// redis
	GetRedis() IRedisOption
}
