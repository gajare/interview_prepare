package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.isitdownrightnow123454321.com",
		"https://www.uptrends.com",
	}
	for _, url := range urls {
		if CheckStatus(url) {
			fmt.Println(url, "Up")
		} else {
			fmt.Println(url, "Down")
		}
	}
}

func CheckStatus(url string) bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	fmt.Println("\t url:", url, "\t status:", resp.StatusCode)
	return resp.StatusCode >= 200 && resp.StatusCode < 400
}
