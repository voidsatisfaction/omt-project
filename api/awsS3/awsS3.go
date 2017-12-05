package awsS3

import (
	"fmt"
	"omt-project/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateAwsSession() (*session.Session, error) {
	c := config.Setting()
	fmt.Printf("awsregion: %+v\n", c)
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(c.AwsRegion),
		Credentials: credentials.NewStaticCredentials(c.AwsAccessKeyID, c.AwsSecretAccessKey, ""),
	})
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func CreateS3Client() (*s3.S3, error) {
	sess, err := CreateAwsSession()
	if err != nil {
		return nil, err
	}
	return s3.New(sess), nil
}

func GetWordKey(userId string) string {
	c := config.Setting()
	return fmt.Sprintf("%s%s", c.AwsS3BucketWordsKey, userId)
}

func GetUserKey(userId string) string {
	c := config.Setting()
	return fmt.Sprintf("%s%s", c.AwsS3BucketUsersKey, userId)
}
