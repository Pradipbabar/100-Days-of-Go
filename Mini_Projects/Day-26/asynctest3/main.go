package main

import (
	"fmt"
	"net/http"
	"time"
)

type urlStatus struct {
	url    string
	status bool
}

//checks and prints a message if a website is up or down
func checkUrl(url string, c chan urlStatus) {
	client := http.Client {
		Timeout: time.Second * 5,
	}
	_, err := client.Get(url)
	if err != nil {
		// The website is down
		c <- urlStatus{url, false}
	} else {
		// The website is up
		c <- urlStatus{url, true}
	}
}

func main() {
	// A slice of sample websites
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}

	c := make(chan urlStatus)
	for _, url := range urls {
		go checkUrl(url, c)

	}
	result := make([]urlStatus, len(urls))
	for i := range result{
		select {
		case res := <-c:
			result[i] = res
		case <-time.After(time.Second * 5):
			fmt.Println("Timeout occurred",urls[i])
			result[i] = urlStatus{urls[i], false}
		}
	}
	
	
	for _, res := range result {
		if res.status {
			fmt.Println(res.url, "is up.")
		} else {
			fmt.Println(res.url, "is down !!")
		}
	}
}



