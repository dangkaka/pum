package main

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"io"
	"net/http"
	"os"
)

func download(obj DownloadObject) {
	resp, err := http.Get(obj.DownloadUrl)
	if err != nil {
		fmt.Println("Could not reach download url", obj.DownloadUrl, err)
		return
	}
	defer resp.Body.Close()

	fullName := fmt.Sprintf("%s.%s", obj.Title, obj.Type)

	out, err := os.Create(fullName)
	defer out.Close()
	if err != nil {
		fmt.Println("Could not create file", err)
		return
	}
	bar := pb.New(int(resp.ContentLength)).SetUnits(pb.U_BYTES)
	bar.ShowSpeed = true
	bar.ShowTimeLeft = true
	bar.ShowBar = true
	bar.ShowPercent = true
	bar.SetWidth(80)
	bar.Prefix(fullName)

	bar.Start()

	rd := bar.NewProxyReader(resp.Body)
	_, err = io.Copy(out, rd)
	if err != nil {
		fmt.Println("Could not copy file", fullName, err)
		return
	}
	bar.Finish()
	return
}
