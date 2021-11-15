package app

type Config struct {
	Title    string
	Desc     string
	Keywords string
}

const (
	title = "Betfate - The revolution Betting & Trading Exchange, Artificial intelligence and Big Data tools , new game changer in sports betting"
	desc  = "BETFATE creates betting & trading platform with elements such as sports betting, trading, social networking, analytics, Big Data and etc"
	key   = "Betfate, Betting, Exchange, Trading, AI, BigData, E-Sports"
)

// Newconfig
func NewConfig() *Config {
	return &Config{
		Title:    title,
		Desc:     desc,
		Keywords: key,
	}
}
