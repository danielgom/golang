package main

import (
	"fmt"
)

func main() {

	/*
		We know that sending and receiving values through a channel is a blocking operation (it accepts values one by one)
		however there is a way to send the values without blocking , this is through a buffered channel
	*/

	ch := make(chan int, 6) // We have a buffered channel that can handle up to 6 values, we can compare this with the last program

	go func() {
		defer close(ch)
		for i := 0; i < 6; i++ {
			fmt.Println("Sending:", i)
			ch <- i
		}
	}()

	/*
		Concurrent (main go routine), the code here is running in the main go routine means that while other go routines
		are working on sending information to the channel this code is running in the main routine
	*/


	fmt.Println("hello")
	fmt.Println("how are you?")

	for v := range ch {
		fmt.Println("Received: ", v)
	}

}
