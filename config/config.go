package config

import "os"

// Config is configuration of app
type Config struct {
	Host                     string
	AppEnv                   string
	Port                     string
	ChannelSecret            string
	ChannelToken             string
	AwsRegion                string
	AwsAccessKeyID           string
	AwsSecretAccessKey       string
	AwsS3Bucket              string
	AwsS3BucketUsersKey      string
	AwsS3BucketWordsKey      string
	AwsS3BucketQuizTimersKey string
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
	c.Host = "https://omt-project.herokuapp.com"
	c.AwsRegion = os.Getenv("AWS_REGION")
	c.AwsAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	c.AwsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	c.AwsS3Bucket = "omt-project"
	c.AwsS3BucketUsersKey = "users/"
	c.AwsS3BucketWordsKey = "words/"
	c.AwsS3BucketQuizTimersKey = "quizTimers/"
	return c
}
