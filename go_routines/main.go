package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	for _, link := range sites {
		go checkLink(link, c)
	}

	for l := range c {
		time.Sleep(3 * time.Second)
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is Up!!.")
	c <- link
}
