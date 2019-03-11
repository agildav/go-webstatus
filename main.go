package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	c := make(chan string)

	links := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.facebook.com",
		"http://www.instagram.com",
		"http://www.tibia.com",
	}

	for _, link := range links {
		go checkLink(link, c)
	}

	for link := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			go checkLink(l, c)
		}(link)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
	return
}
