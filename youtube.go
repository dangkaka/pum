package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const linkDownload = "http://www.youtube.com/get_video_info?&video_id="

type Youtube struct {
	Url string
}

type YoutubeResponse struct {
	Author      string
	DownloadURL string
	Title       string
	Formats     []Format
}

// Format ...
type Format struct {
	VideoType, Quality, URL string
}

func NewYoutubeHandler(url string) *Youtube {
	return &Youtube{url}
}

func (y *Youtube) Get() (*YoutubeResponse, error) {
	video := &YoutubeResponse{}

	if y.Url == "" {
		return nil, errors.New("Empty Url")
	}

	urlList := strings.Split(y.Url, "/")
	if len(urlList) < 4 {
		return nil, errors.New("Invalid Url")
	}

	//get videoId from watch?v=U1M5GDNNhCo
	videoID := urlList[3][8:]

	resp, err := http.Get(linkDownload + videoID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	queryString, _ := ioutil.ReadAll(resp.Body)

	u, _ := url.Parse("?" + string(queryString))
	query := u.Query()

	video.Title = query.Get("title")
	video.Author = query.Get("author")

	formatParam := strings.Split(query.Get("url_encoded_fmt_stream_map"), ",")
	var formats []Format
	for _, f := range formatParam {
		furl, _ := url.Parse("?" + f)
		fquery := furl.Query()
		formats = append(
			formats,
			Format{
				VideoType: fquery.Get("type"),
				Quality:   fquery.Get("quality"),
				URL:       fquery.Get("url"),
			})
	}
	video.Formats = formats
	video.DownloadURL = formats[0].URL //best

	return video, nil
}
