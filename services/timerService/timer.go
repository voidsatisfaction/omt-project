package timerService

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

func NewQuizTimerMap() quizTimerMap {
	return make(quizTimerMap)
}

// This create new Timer
func CreateQuizTimer(timerId string) error {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return err
	}

	qtm := NewQuizTimerMap()

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

func ReadAllIdsByTimerId(timerId string) ([]string, error) {
	qtm, err := GetQuizTimerMap(timerId)
	if err != nil {
		return nil, err
	}

	var quizTimerIds []string
	for userId, _ := range qtm {
		quizTimerIds = append(quizTimerIds, userId)
	}

	return quizTimerIds, nil
}

// ex. timerId 2:34
func GetQuizTimerMap(timerId string) (quizTimerMap, error) {
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

	qtm := NewQuizTimerMap()
	if err := json.Unmarshal(body, &qtm); err != nil {
		return nil, err
	}

	return qtm, nil
}

func AddQuizTimer(userId, timerId string) error {
	qtm, err := GetQuizTimerMap(timerId)
	if err != nil {
		// TODO: check error code 404. This code is not elegant
		if strings.Contains(err.Error(), "code: 404") {
			CreateQuizTimer(timerId)
			qtm = NewQuizTimerMap()
		} else {
			return err
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

func ExistQuizTimer(timerId string) (bool, error) {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return false, err
	}

	timerKey := awsS3.GetTimerKey(timerId)
	_, err = svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(timerKey),
	})

	if err != nil {
		if strings.Contains(err.Error(), "status code: 404") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
