package yelp

import "os"

const ConfigUrl = "https://api.yelp.com/v3"

type Config struct {
	Url      string
	ApiKey   string
	ClientID string
}

func ConfigFromEnv() Config {
	return Config{
		Url:      ConfigUrl,
		ApiKey:   os.Getenv("YELP_API_KEY"),
		ClientID: os.Getenv("YELP_CLIENT_ID"),
	}
}
