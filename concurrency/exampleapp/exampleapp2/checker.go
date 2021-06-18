package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string, len(links))

	for {
		for x := range links {
			//go checkLink(links[x], c)
			go func(link string, c chan<- string) {
				_, err := http.Get(link)
				if err != nil {
					c <- link + " may be down"
				}
				c <- link + " is up"
			}(links[x], c)
		}

		for i := 0; i < cap(c); i++ {
			fmt.Println(<-c)
		}
		time.Sleep(time.Second * 5)
	}
}

func checkLink(link string, c chan<- string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " might be down"
	}
	c <- link + " is up"
}
