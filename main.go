package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	content, err := fetchURL("https://www.google.com")
	if err != nil {
		panic(err)

	}
	fmt.Print(content)
}
func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
