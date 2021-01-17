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

	// To communicate between go routines we need to create channels
	c := make(chan string) // Creating a new channel of type string, it means that this channel can only transport string type things

	for _, link := range links {
		go checkLink(link, c) // New go routine
	}

	/*
		Channel sending and retrieving messages is a blocking action, everytime we ask for the message we retrieve only one result
		instead of the full 5 results we want (5 results means that we have only 5 links in the slice),
	*/

	// Following code is not the best code for retrieving messages from channel
	/*
		fmt.Println(<- c) // Retrieving message from the channel
		fmt.Println(<- c) // Retrieving second message from the channel
		fmt.Println(<- c) // Retrieving third message from the channel
	*/

	for l := range c {
		/*
			go checkLink(l, c) // This is not going to iterate many times a second, instead it will run the for loop as soon as a message
			// is published to the channel
			time.Sleep(time.Duration(2) * time.Second)
			time.Sleep(time.Second)

		*/
		// Using function literals
		go func(link string) { // Name of the argument passed is link
			time.Sleep(time.Second * 5)
			checkLink(link, c) // Use the link variable
		}(l) // This means that we are passing l as an argument to the function literal
	}
}

func checkLink(link string, c chan string) { // We need to add the parameter channel of type string
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		c <- link // Sending a message to a channel
		return
	}
	fmt.Println(link + " is up!")
	c <- link // Sending a message to a channel
}
