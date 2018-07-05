package main

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func download(obj DownloadObject) {
	resp, err := http.Get(obj.DownloadUrl)
	if err != nil {
		log.Println("Could not reach download url", obj.DownloadUrl)
		return
	}
	defer resp.Body.Close()

	fullName := fmt.Sprintf("%s.%s", obj.Title, obj.Type)

	out, err := os.Create(fullName)
	defer out.Close()
	if err != nil {
		log.Println("Could not create file")
		return
	}
	bar := pb.New(int(resp.ContentLength)).SetUnits(pb.U_BYTES)
	bar.SetRefreshRate(200 * time.Microsecond)
	bar.ShowSpeed = true
	bar.ShowTimeLeft = true
	bar.ShowBar = true
	bar.ShowPercent = true
	bar.Prefix(fullName)

	bar.Start()

	rd := bar.NewProxyReader(resp.Body)
	_, err = io.Copy(out, rd)
	if err != nil {
		log.Println("Could not copy file", fullName)
		return
	}
	bar.Finish()
	return
}
