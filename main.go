package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.facebook.com",
		"http://www.instagram.com",
		"http://www.tibia.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down")
		return
	}

	fmt.Println(link, "is up")
	return
}
