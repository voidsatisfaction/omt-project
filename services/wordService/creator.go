package wordService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"omt-project/api/awsS3"
	"omt-project/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type WordsInfo struct {
	Words []Word
}

type Word struct {
	Name     string
	Meaning  []string
	Priority int
}

func CreateNewWord(userId string) {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return
	}

	wordsInfo := WordsInfo{
		Words: []Word{},
	}

	wordsInfoJson, err := json.Marshal(wordsInfo)
	if err != nil {
		return
	}

	wordKey := fmt.Sprintf("%s%s", c.AwsS3BucketWordsKey, userId)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.AwsS3Bucket),
		Key:    aws.String(wordKey),
		// TODO: casting directly []byte() is not efficient refer: https://qiita.com/ikawaha/items/3c3994559dfeffb9f8c9
		Body: bytes.NewReader(wordsInfoJson),
	})

	if err != nil {
		panic(err)
	}
}
