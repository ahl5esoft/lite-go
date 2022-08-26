package config

import "github.com/ahl5esoft/lite-go/model/contract"

// 默认配置
type Default struct {
	// Mongo
	Mongo string
	// 名字
	Name string
	// 端口
	Port int
	// redis
	Redis DefaultRedis
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
