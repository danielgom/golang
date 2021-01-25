package main

import "fmt"

func main() {

	ch := make(chan int)
	// This makes a FIFO
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println("sending:", i)
			ch <- i
		}
		close(ch) // If we do not close the channel the next for keeps waiting for a message
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
