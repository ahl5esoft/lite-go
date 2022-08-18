package config

// 默认配置
type Default struct {
	// Mongo
	Mongo string
	// 名字
	Name string `yaml:"name"`
	// 端口
	Port int `yaml:"port"`
	// redis
	Redis struct {
		Host     string
		Password string
	}
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

func (m Default) GetRedis() struct {
	Host     string
	Password string
} {
	return m.Redis
}
