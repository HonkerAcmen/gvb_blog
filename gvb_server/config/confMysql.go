package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	Config   string `yaml:"config"` // 高级配置，例如charset=utf-8
	User     string `yaml:"user"`
	Passwd   string `yaml:"passwd"`
	LogLevel string `yaml:"log_level"`
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Passwd + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
