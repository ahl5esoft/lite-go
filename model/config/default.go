package config

import "github.com/ahl5esoft/lite-go/model/contract"

// 默认配置
type Default struct {
	// 认证密钥
	AuthCipher string
	// 网关地址
	Gateway string
	// Mongo
	Mongo string
	// 名字
	Name string
	// 端口
	Port int
	// redis
	Redis DefaultRedis
}

func (m Default) GetAuthCipher() string {
	return m.AuthCipher
}

func (m Default) GetGateway() string {
	return m.Gateway
}

func (m Default) GetMongo() string {
	return m.Mongo
}

func (m Default) GetName() string {
	return m.Name
}

func (m Default) GetPort() int {
	return m.Port
}

func (m Default) GetRedis() contract.IRedisOption {
	return m.Redis
}
