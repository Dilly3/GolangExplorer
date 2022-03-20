package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golong.org",
		"https://amazon.com",
	}

	for _, link := range links {

		_, err := http.Get(link)
		HandleError(err)

	}
}

func CheckLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		HandleError(err)
		return
	}
	fmt.Println("Link is up", link)

}
func HandleError(err error) {
	if err != nil {
		fmt.Println("cant get to link")
	}

}
