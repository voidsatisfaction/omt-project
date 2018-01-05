package userService

import (
	"omt-project/api/awsS3"
	"omt-project/config"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func IsUserExist(userId string) (bool, error) {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return false, err
	}

	userKey := awsS3.GetUserKey(userId)
	_, err = svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(userKey),
	})
	if err != nil {
		if strings.Contains(err.Error(), "status code: 404") {
			return false, err
		}
	}
	return true, nil
}
