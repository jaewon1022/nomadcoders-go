package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errHttpRequestFailed = errors.New("http request failed")

func main() {
	var results = make(map[string]string)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"http://asdf",
	}

	for _, url := range urls {
		err := hitURL(url)

		result := "OK"

		if err != nil {
			result = "FAILED"
		}

		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking URL :", url)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		return errHttpRequestFailed
	} 

	return nil

}