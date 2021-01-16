package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://google.org",
		"http://amazon.com",
	}

	// To communicate between go routines we need to create channels
	c := make(chan string) // Creating a new channel of type string, it means that this channel can only transport string type things

	for _, link := range links {
		go checkLink(link, c)  // New go routine
	}
}

func checkLink(link string, c chan string) { // We need to add the parameter channel of type string
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		return
	}

	fmt.Println(link + " is up!")
}
