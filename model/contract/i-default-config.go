package contract

// 默认配置
type IDefaultConfig interface {
	GetGateway() string
	GetMongo() string
	GetName() string
	GetPort() int
	GetRedis() IRedisOption
}
