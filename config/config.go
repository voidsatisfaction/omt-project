package config

import "os"

// Config is configuration of app
type Config struct {
	AppEnv        string
	Port          string
	ChannelSecret string
	ChannelToken  string
}

// Setting function returns configuration
func Setting() *Config {
	c := &Config{}
	switch os.Getenv("APP_ENV") {
	case "DEV":
		c.AppEnv = "DEV"
		c.Port = "9001"
		c.ChannelSecret = os.Getenv("CHANNEL_SECRET")
		c.ChannelToken = os.Getenv("CHANNEL_TOKEN")
	case "PROD":
		c.AppEnv = "PROD"
		c.Port = os.Getenv("PORT")
		c.ChannelSecret = os.Getenv("CHANNEL_SECRET")
		c.ChannelToken = os.Getenv("CHANNEL_TOKEN")
	}
	return c
}
