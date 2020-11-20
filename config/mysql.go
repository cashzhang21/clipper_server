package config

import "time"

type EndpointConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type DSNConfig struct {
	UserName string          `yaml:"userName"`
	Password string          `yaml:"password"`
	Endpoint *EndpointConfig `yaml:"endpoint"`
	DBName   string          `yaml:"dbName"`
	Options  []string        `yaml:"options"`
}

type MysqlConfig struct {
	// DSN配置
	DSN *DSNConfig `yaml:"dsn"`
	// 最大可用数量
	Active int `yaml:"active"`
	// 最大闲置数量
	Idle int `yaml:"idle"`
	// 闲置超时时间
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}
