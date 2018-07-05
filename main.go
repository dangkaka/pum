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

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Too few arguments")
		fmt.Println("Usage: pum <URLs>...")
		return
	}
	for _, videoUrl := range args {
		switch {
		case strings.Contains(videoUrl, youtube):
			youtube := NewYoutubeHandler(videoUrl)
			downloadUrl, err := youtube.Get()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(downloadUrl)
		case strings.Contains(videoUrl, zing):
			zing := NewZingHandler(videoUrl)
			downloadUrl, err := zing.GetBest()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(downloadUrl)
		case strings.Contains(videoUrl, nct):
			nct := NewNCTHandler(videoUrl)
			nct.Get()
		}
	}
}
