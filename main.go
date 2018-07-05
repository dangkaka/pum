package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

const (
	nct     = "nhaccuatui"
	zing    = "zing"
	youtube = "youtube"
)

type DownloadObject struct {
	Url         string
	Author      string
	Title       string
	DownloadUrl string
	Type        string
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Too few arguments")
		fmt.Println("Usage: pum <URLs>...")
		return
	}
	for _, url := range args {
		switch {
		case strings.Contains(url, youtube):
			youtube := NewYoutubeHandler(url)
			downloadObject, err := youtube.GetDownloadObject()
			if err != nil {
				log.Fatal(err)
			}
			download(*downloadObject)
		case strings.Contains(url, zing):
			zing := NewZingHandler(url)
			downloadObject, err := zing.GetDownloadObject()
			if err != nil {
				log.Fatal(err)
			}
			download(*downloadObject)
		case strings.Contains(url, nct):
			nct := NewNCTHandler(url)
			nct.GetDownloadObject()
		}
	}
	fmt.Println("Done!")
}
