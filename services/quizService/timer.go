package quizService

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"omt-project/api/awsS3"
	"omt-project/config"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type quizTimerMap map[string]*json.RawMessage

// This create new Timer
func CreateQuizTimer(timerId string) error {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return err
	}

	qtm := make(quizTimerMap)

	qtmJSON, err := json.Marshal(qtm)
	if err != nil {
		return err
	}

	timerKey := awsS3.GetTimerKey(timerId)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(timerKey),
		Body:   bytes.NewReader(qtmJSON),
	})
	if err != nil {
		return err
	}

	return nil
}

func ReadQuizTimer(timerId string) (quizTimerMap, error) {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return nil, err
	}

	timerKey := awsS3.GetTimerKey(timerId)
	res, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(timerKey),
	})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	qtm := quizTimerMap{}
	if err := json.Unmarshal(body, &qtm); err != nil {
		return nil, err
	}

	return qtm, nil
}

func AddQuizTimer(userId, timerId string) error {
	qtm, err := ReadQuizTimer(timerId)
	if err != nil {
		// TODO: check error code 404. This code is not elegant
		if strings.Contains(err.Error(), "code: 404") {
			CreateQuizTimer(timerId)
			qtm = quizTimerMap{}
		}
	}
	if _, ok := qtm[userId]; ok {
		return nil
	}
	qtm[userId] = &json.RawMessage{'1'}

	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return err
	}

	qtmJSON, err := json.Marshal(qtm)
	if err != nil {
		return err
	}

	timerKey := awsS3.GetTimerKey(timerId)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(timerKey),
		Body:   bytes.NewReader(qtmJSON),
	})
	if err != nil {
		return err
	}

	return nil
}
