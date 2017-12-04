package glosbe

import (
	"fmt"
	"net/http"
)

const (
	// GlosbeHost = "https://glosbe.com/gapi/translate?from=eng&dest=kor&format=json&pretty=true&phrase=dragon"
	GlosbeHost = "https://glosbe.com"
)

// GlosbeClient
type GlosbeClient struct {
	http.Client
	host string
}

func CreateGlosbeClient() *GlosbeClient {
	gc := &GlosbeClient{}
	gc.host = GlosbeHost
	return gc
}

type GlosbeParameter struct {
	LanguageFrom string
	LanguageTo   string
	Phrase       string
}

// CreateGlosbeTranslateRequest is making *http.Requst to translate word
func CreateGlosbeTranslateRequest(gp *GlosbeParameter) (*http.Request, error) {
	lFrom := gp.LanguageFrom
	lTo := gp.LanguageTo
	phrase := gp.Phrase

	reqURL := fmt.Sprintf(
		"%s/gapi/translate?from=%s&dest=%s&format=json&pretty=true&phrase=%s",
		GlosbeHost,
		lFrom,
		lTo,
		phrase,
	)
	req, err := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36")
	req.Header.Add("Accept", "*/*")
	if err != nil {
		return nil, err
	}
	return req, nil
}

type GlosbeResponse struct {
	Result string `json:"result"`
	Tucs   []Tuc  `json:"tuc"`
}

type Tuc struct {
	Phrase Phrase
}

type Phrase struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}

func ExtractTenMeaning(gRes *GlosbeResponse) []string {
	var s []string
	check := make(map[string]bool)
	for i, tuc := range gRes.Tucs {
		if i > 9 {
			break
		}
		text := tuc.Phrase.Text
		if _, exist := check[text]; !exist {
			check[text] = true
			s = append(s, text)
		}
	}

	return s
}
