package wordService

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"omt-project/api/awsS3"
	"omt-project/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type WordsInfo struct {
	Words Words `json:"words"`
}

type Words []Word

type Word struct {
	Name     string   `json:"name"`
	Meaning  []string `json:"meaning"`
	Priority int      `json:"priority"`
}

func CreateNewWord(userId string) error {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return err
	}

	wordsInfo := WordsInfo{
		Words: Words{},
	}

	wordsInfoJson, err := json.Marshal(wordsInfo)
	if err != nil {
		return err
	}

	wordKey := awsS3.GetWordKey(userId)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(wordKey),
		// TODO: casting directly []byte() is not efficient refer: https://qiita.com/ikawaha/items/3c3994559dfeffb9f8c9
		Body: bytes.NewReader(wordsInfoJson),
	})
	if err != nil {
		return err
	}

	return nil
}

func ReadWord(userId string) (*WordsInfo, error) {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return nil, err
	}

	wordKey := awsS3.GetWordKey(userId)
	res, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(wordKey),
	})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	wi := &WordsInfo{}
	if err := json.Unmarshal(body, wi); err != nil {
		return nil, err
	}

	return wi, nil
}

func Addword(userId string, wordName string, meaning []string) error {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return err
	}

	wordsInfo, err := ReadWord(userId)
	if err != nil {
		return err
	}

	w := Word{
		Name:     wordName,
		Meaning:  meaning,
		Priority: 100,
	}
	wordsInfo.Words = append(wordsInfo.Words, w)

	wordsInfoJSON, err := json.Marshal(wordsInfo)
	if err != nil {
		return err
	}

	wordKey := awsS3.GetWordKey(userId)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(wordKey),
		// TODO: casting directly []byte() is not efficient refer: https://qiita.com/ikawaha/items/3c3994559dfeffb9f8c9
		Body: bytes.NewReader(wordsInfoJSON),
	})
	if err != nil {
		return err
	}

	return nil
}
