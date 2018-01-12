package wordService

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"omt-project/config"
	"omt-project/src/api/awsS3"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type WordsInfo struct {
	Words Words `json:"words"`
}

func NewWordsInfo() *WordsInfo {
	return &WordsInfo{}
}

func (wi *WordsInfo) AssignWords(ws Words) {
	wi.Words = ws
}

func (wi *WordsInfo) addWord(w Word) {
	wi.Words = append(wi.Words, w)
}

func (wi *WordsInfo) findWordIndex(wn string) int {
	for i, w := range wi.Words {
		if wn == w.Name {
			return i
		}
	}
	return -1
}

func (wi *WordsInfo) existWord(wn string) bool {
	for _, w := range wi.Words {
		if wn == w.Name {
			return true
		}
	}
	return false
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

	// TODO: move this logic to s3 services
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

func ReadWords(userId string) (*WordsInfo, error) {
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

func UpdateNewWord(userId string, wi *WordsInfo) error {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return err
	}

	wordsInfoJson, err := json.Marshal(wi)
	if err != nil {
		return err
	}

	// TODO: move this logic to s3 services
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

func Addword(userId string, wordName string, meaning []string) error {
	c := config.Setting()
	svc, err := awsS3.CreateS3Client()
	if err != nil {
		return err
	}

	wordsInfo, err := ReadWords(userId)
	if err != nil {
		return err
	}

	if wordsInfo.existWord(wordName) {
		i := wordsInfo.findWordIndex(wordName)
		wordsInfo.Words[i].Priority = 100
	} else {
		w := Word{
			Name:     wordName,
			Meaning:  meaning,
			Priority: 100,
		}
		wordsInfo.addWord(w)
	}

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
