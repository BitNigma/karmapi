package app

type Config struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

// Newconfig
func NewConfig() *Config {
	return &Config{}
}
