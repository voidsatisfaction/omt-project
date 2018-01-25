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
	// common setting
	c.ChannelSecret = os.Getenv("CHANNEL_SECRET")
	c.ChannelToken = os.Getenv("CHANNEL_TOKEN")
	c.AwsRegion = os.Getenv("AWS_REGION")
	c.AwsAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	c.AwsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	c.AwsS3Bucket = "omt-project"
	c.AwsS3BucketUsersKey = "users/"
	c.AwsS3BucketWordsKey = "words/"
	c.AwsS3BucketQuizTimersKey = "quizTimers/"

	// environment dependent setting
	switch os.Getenv("APP_ENV") {
	case "DEV":
		c.AppEnv = "DEV"
		c.Port = "9001"
		c.Host = "http://localhost:19000"
	case "STAGING":
		// TODO: add staging setting
	case "PROD":
		c.AppEnv = "PROD"
		c.Port = "9000"
		c.Host = "https://hiyoko-teacher.let.media.kyoto-u.ac.jp/live"
	}
	return c
}
