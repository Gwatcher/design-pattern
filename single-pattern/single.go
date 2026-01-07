package single_pattern

import "sync"

type Config struct {
	Name    string
	Ip      string
	Port    string
	Address string
}

var (
	config *Config
	once   sync.Once
)

func ConfigInit() {
	config = &Config{
		Name: "ljl",
	}
}

func GetConfig() *Config {
	once.Do(func() {
		ConfigInit()
	})
	return config
}
