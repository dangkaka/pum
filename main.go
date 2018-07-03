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
		youtubeResponse, err := Get(videoUrl)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(youtubeResponse.DownloadURL)
	}
}
