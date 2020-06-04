package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for link := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(link)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is broken")
		c <- link
		return
	}
	fmt.Println(link, "is OK")
	c <- link
}
