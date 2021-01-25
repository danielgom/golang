package main

import "fmt"

func main() {

	// We can specify in a function whether the channel can either only receive or send messages
	ch1 := make(chan string)
	ch2 := make(chan string)
	go genMsg(ch1, "message")
	go relayMsg(ch1, ch2)
	v := <-ch2
	fmt.Println(v)

}

func genMsg(ch1 chan<- string, msg string) { // This syntax means that ch1 is only available for receiving
	ch1 <- msg
	//v := <- ch1 invalid since ch1 here is only available for receiving messages
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) { // Ch1 only sends messages, Ch2 only receives messages
	m := <-ch1
	ch2 <- m
}
