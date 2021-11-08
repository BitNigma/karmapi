package app

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

// Newconfig
func NewConfig(path string) *Config {

	vi := viper.New()
	vi.SetConfigFile(path)
	if err := vi.ReadInConfig(); err != nil {
		log.Fatal("can't read config file")
	}

	return &Config{
		Port: vi.GetString("port"),
		Host: vi.GetString("addr"),
	}
}
