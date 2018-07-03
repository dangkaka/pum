package main

import (
	"flag"
	"fmt"
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
		youtube := NewYoutubeHandler(videoUrl)
		response, err := youtube.Get()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(response.Title, response.Author, response.DownloadURL)
	}
}
