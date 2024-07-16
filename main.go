package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errHttpRequestFailed = errors.New("http request failed")

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}

	for _, url := range urls {
		hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking URL...")
	
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		return errHttpRequestFailed
	} 

	return nil

}