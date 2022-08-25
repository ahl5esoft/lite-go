package contract

// 默认配置
type IDefaultConfig interface {
	GetMongo() string
	GetName() string
	GetPort() int
	GetRedis() IRedisOption
}
