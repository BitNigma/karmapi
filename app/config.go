package app

type Config struct {
	Title    string
	Desc     string
	Keywords string
}

const (
	title = "TaroPI - Prection Ecosystem"
	desc  = "TaroPI make your Dream reality, Prediction Taro cards, NFT , Crypto Token"
	key   = "TaroPI, Taro, Crypto, Token, Precition, Fate"
)

// Newconfig
func NewConfig() *Config {
	return &Config{
		Title:    title,
		Desc:     desc,
		Keywords: key,
	}
}
