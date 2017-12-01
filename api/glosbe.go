package api

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
	Phrase struct {
		Text     string `json:"text"`
		Language string `json:"language"`
	}
}
