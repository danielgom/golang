package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		Timeouts are kind of important when we want to connect to external resources or that otherwise need to bound
		execution time, we can implement timeouts easily with channels and the select keyword
	*/

	// In the following examples we are going to show how timeouts work

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2) // simulating a 2 second process
		c1 <- "result 1"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1): // timeout, if 1 second passes and no response from channel 1 is found
		fmt.Println("timeout on channel 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1) // simulating a 1 second process
		c2 <- "result 2"
	}()
	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(time.Second * 2): // timeout, if 2 second passes and no response from channel 2 is found
		fmt.Println("timeout on channel 2")
	}

	// Result : timeout on channel 1 and result 2 on channel 2

	// Non-Blocking operations, following examples are going to show how non-blocking select works with a default clause
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg: // Since the channel is not buffered, this message is not sent
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
