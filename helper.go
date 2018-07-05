package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
)

func parseResponse(input []byte, pattern string) (string, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	submatches := r.FindSubmatch(input)
	if len(submatches) != 2 {
		return "", errors.New("Failed to parse data")
	}
	return string(submatches[1]), nil
}

func GetBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	respBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}
