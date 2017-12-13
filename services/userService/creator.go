package userService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"omt-project/api/awsS3"
	"omt-project/config"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type UserInfo struct {
	Id                    string      `json:"id"`
	ConsecutiveActionDays int         `json:"consecutive_action_days"`
	LastActionTime        time.Time   `json:"last_action_time"`
	PushTimes             []time.Time `json:"push_times"`
}

func CreateNewUser(userId string) {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return
	}

	userInfo := UserInfo{
		Id: userId,
		ConsecutiveActionDays: 0,
		LastActionTime:        time.Now(),
		PushTimes:             []time.Time{},
	}

	userInfoJson, err := json.Marshal(userInfo)
	if err != nil {
		return
	}

	userKey := fmt.Sprintf("%s%s", c.AwsS3BucketUsersKey, userId)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(userKey),
		// TODO: casting directly []byte() is not efficient refer: https://qiita.com/ikawaha/items/3c3994559dfeffb9f8c9
		Body: bytes.NewReader(userInfoJson),
	})

	if err != nil {
		panic(err)
	}
}
