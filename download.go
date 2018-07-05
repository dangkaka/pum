package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/cheggaaa/pb"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type PreparedDownloadObject struct {
	Resp *http.Response
	Name string
}

func download(objects []DownloadObject) {
	objs, err := prepare(objects)
	if err != nil {
		log.Println("failed to create objects response")
		return
	}
	write(objs)
	return
}

func write(objs []PreparedDownloadObject) {

	var wg sync.WaitGroup
	wg.Add(len(objs))

	var barList []*pb.ProgressBar
	for _, o := range objs {
		bar := pb.New(int(o.Resp.ContentLength)).SetUnits(pb.U_BYTES)
		bar.ShowSpeed = true
		bar.ShowTimeLeft = true
		bar.ShowBar = true
		bar.ShowPercent = true
		bar.Prefix(o.Name)

		barList = append(barList, bar)
	}

	pool, err := pb.StartPool(barList...)
	if err != nil {
		logrus.WithError(err).Error("cannot start pool")
		return
	}

	for i, v := range objs {
		go func(o PreparedDownloadObject) {
			defer wg.Done()
			defer o.Resp.Body.Close()

			rd := barList[i].NewProxyReader(o.Resp.Body)

			out, err := os.OpenFile(o.Name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			defer out.Close()
			if err != nil {
				logrus.WithError(err).Errorf("cannot open file, file name = %s", o.Name)
				return
			}

			_, err = io.Copy(out, rd)
			if err != nil {
				logrus.WithError(err).Errorf("cannot copy file, file name = %s", o.Name)
				return
			}

			time.Sleep(500 * time.Millisecond)
		}(v)
	}
	pool.Stop()
	wg.Wait()
}

func prepare(objects []DownloadObject) ([]PreparedDownloadObject, error) {
	var objs []PreparedDownloadObject
	var wg sync.WaitGroup
	wg.Add(len(objects))

	for _, v := range objects {
		go func(obj DownloadObject) {
			defer wg.Done()
			resp, err := http.Get(obj.DownloadUrl)
			if err != nil {
				logrus.WithError(err).Error("failed to get response from download url")
				return
			}

			fullName := fmt.Sprintf("%s.%s", obj.Title, obj.Type)

			out, err := os.Create(fullName)
			defer out.Close()
			if err != nil {
				log.Println("Could not create file")
				return
			}
			objs = append(objs, PreparedDownloadObject{
				resp,
				fullName,
			})
		}(v)
	}
	wg.Wait()

	return objs, nil
}
