package app

import (
	"crypto/tls"
)

type Config struct {
	Title    string
	Desc     string
	Keywords string
	Tl       *tls.Config
}

const (
	title = "KARMAPI, Shape Your Future with AI and Prediction"
	desc  = "KARMAPI - Ecosystem, Decentralised Prediction, Fortune telling"
	key   = "KARMAVERSE, SOCIALFI, GAMEFI, Decentralised Prediction, Fortune telling , FATE AND FUTURE"
)

// Newconfig
func NewConfig() *Config {

	/* // load tls certificates
	TLSCert, err := tls.LoadX509KeyPair(CACertFilePath, KeyFilePath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
		return nil
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{TLSCert},
	} */

	return &Config{
		Title:    title,
		Desc:     desc,
		Keywords: key,
		//Tl:       tlsConfig,
	}
}
