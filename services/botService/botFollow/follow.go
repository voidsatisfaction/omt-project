package botFollow

import (
	"fmt"
	"omt-project/api/awsS3"
	"omt-project/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateNewUser(userId string) {
	fmt.Println(userId)
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return
	}

	key := fmt.Sprintf("%s%s", c.AwsS3BucketUsersKey, userId)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(svc)
}
