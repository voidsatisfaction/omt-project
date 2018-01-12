package userService

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"omt-project/config"
	"omt-project/src/api/awsS3"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type UserInfo struct {
	Id                    string    `json:"id"`
	ConsecutiveActionDays int       `json:"consecutive_action_days"`
	LastActionTime        time.Time `json:"last_action_time"`
	PushTimes             []string  `json:"push_times"`
}

func NewUserInfo(id string, cad int, lat time.Time, pt []string) *UserInfo {
	return &UserInfo{
		id, cad, lat, pt,
	}
}

func CreateNewUser(userId string) error {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return err
	}

	userInfo := UserInfo{
		Id: userId,
		ConsecutiveActionDays: 0,
		LastActionTime:        time.Now(),
		PushTimes:             []string{},
	}

	userInfoJson, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}

	userKey := awsS3.GetUserKey(userId)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(userKey),
		// TODO: casting directly []byte() is not efficient refer: https://qiita.com/ikawaha/items/3c3994559dfeffb9f8c9
		Body: bytes.NewReader(userInfoJson),
	})

	if err != nil {
		return err
	}
	return nil
}

func ReadUserInfo(userId string) (*UserInfo, error) {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return nil, err
	}

	userKey := awsS3.GetUserKey(userId)
	res, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(userKey),
	})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	ui := &UserInfo{}
	if err := json.Unmarshal(body, ui); err != nil {
		return nil, err
	}

	return ui, nil
}

func UpdateUserInfo(userId string, ui *UserInfo) error {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return err
	}

	newUserInfoJson, err := json.Marshal(ui)
	if err != nil {
		return err
	}

	userKey := awsS3.GetUserKey(userId)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(userKey),
		// TODO: casting directly []byte() is not efficient refer: https://qiita.com/ikawaha/items/3c3994559dfeffb9f8c9
		Body: bytes.NewReader(newUserInfoJson),
	})

	if err != nil {
		return err
	}
	return nil
}
