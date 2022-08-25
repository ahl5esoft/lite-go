package config

import "fmt"

// 默认Redis
type DefaultRedis struct {
	Port     int
	Host     string
	Password string
}

func (m DefaultRedis) GetAddr() string {
	if m.Port == 0 {
		m.Port = 6379
	}

	return fmt.Sprintf("%s:%d", m.Host, m.Port)
}

func (m DefaultRedis) GetPassword() string {
	return m.Password
}
