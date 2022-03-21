package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var c = make(chan string)
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	for _, link := range links {

		go CheckLink(link, c)

	}
	for li := range c { // wait on a channel and assign any message that come in to link
		go func(s string) {
			time.Sleep(2 * time.Second)
			CheckLink(s, c)
		}(li)
	}

}

func CheckLink(link string, ch chan string) {

	_, err := http.Get(link)
	if err != nil {
		HandleError(err)
		ch <- "Link is down"
	}
	fmt.Println("link is up", link)
	ch <- link
	return

}
func HandleError(err error) {
	if err != nil {
		fmt.Println("cant get to link")
	}

}
