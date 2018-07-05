package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

const (
	nctWmUrl    = "https://m.nhaccuatui.com/bai-hat/%s"
	nctLinkInfo = "https://m.nhaccuatui.com/ajax/get-media-info?key1=%s&key2=&key3="
)

type NCT struct {
	Url string
}

type NCTResponse struct {
	Data struct {
		Title       string `json:"title"`
		Author      string `json:"singerTitle"`
		DownloadUrl string `json:"location"`
	} `json:"data"`
}

func NewNCTHandler(url string) *NCT {
	return &NCT{url}
}

func (nct *NCT) GetDownloadObject() (*DownloadObject, error) {
	response, err := nct.Parse()
	if err != nil {
		return nil, err
	}
	return &DownloadObject{
		Url:         nct.Url,
		Title:       response.Data.Title,
		Author:      response.Data.Author,
		DownloadUrl: response.Data.DownloadUrl,
		Type:        "mp3",
	}, nil
}

func (nct *NCT) Parse() (*NCTResponse, error) {
	response := &NCTResponse{}
	if nct.Url == "" {
		return nil, errors.New("Empty Url")
	}

	urlList := strings.Split(nct.Url, "/")
	if len(urlList) < 5 {
		return nil, errors.New("Invalid Url")
	}

	videoID := urlList[4][:]

	respBytes, err := GetBody(fmt.Sprintf(nctWmUrl, videoID))

	keyPattern := "songencryptkey=\"([a-zA-Z0-9]*)\""

	key, err := parseResponse(respBytes, keyPattern)
	fmt.Println(fmt.Sprintf(nctWmUrl, videoID))
	if err != nil {
		return nil, err
	}

	downloadSourceUrl := fmt.Sprintf(nctLinkInfo, key)

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
