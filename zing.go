package main

import (
	"errors"
	"fmt"
	"encoding/json"
	"net/http"
)

const linkDownloadSong = "https://mp3.zing.vn/xhr/media/get-url-download?type=audio&sig=%s&code=%s"

type Zing struct {
	Url string
}

type ZingResponse struct {
	Data struct {
		Title  string `json:"title"`
		Artist string `json:"artist"`
		Sources struct {
			Url  Source `json:"128"`
			Url2 Source `json:"320"`
			Url3 Source `json:"lossless"`
		} `json:"sources"`
	} `json:"data"`
}

type Source struct {
	Link string `json:"link"`
}

func NewZingHandler(url string) *Zing {
	return &Zing{url}
}

func (zing *Zing) Get() (*ZingResponse, error) {
	response := &ZingResponse{}
	if zing.Url == "" {
		return nil, errors.New("Empty Url")
	}

	respBytes, err := GetBody(zing.Url)

	sigPattern := "data-sig=\"([a-zA-Z0-9]*)\""

	sig, err := parseResponse(respBytes, sigPattern)
	if err != nil {
		return nil, err
	}

	codePattern := "data-code=\"([a-zA-Z0-9]*)\""
	code, err := parseResponse(respBytes, codePattern)
	if err != nil {
		return nil, err
	}

	downloadSourceUrl := fmt.Sprintf(linkDownloadSong, sig, code)

	res, err := http.Get(downloadSourceUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

