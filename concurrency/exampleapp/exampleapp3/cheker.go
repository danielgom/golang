package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	// Same example with waitGroups

	var wg sync.WaitGroup

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for {
		for x := range links {
			wg.Add(1)
			go checkLink(links[x], &wg)
		}
		wg.Wait()
		time.Sleep(time.Second * 5)
	}
}

func checkLink(link string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down")
	}
	fmt.Println(link + " is up")
}
