package app

type Config struct {
	Title    string
	Desc     string
	Keywords string
}

const (
	title = "KARMAPI, Shape Your Future with AI and Prediction"
	desc  = "KARMAPI - Ecosystem, Decentralised Prediction, Fortune telling"
	key   = "KARMAVERSE, SOCIALFI, GAMEFI, Decentralised Prediction, Fortune telling , FATE AND FUTURE"
)

// Newconfig
func NewConfig() *Config {
	return &Config{
		Title:    title,
		Desc:     desc,
		Keywords: key,
	}
}
