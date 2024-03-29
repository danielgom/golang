package main

import "fmt"

func main() {

	// Channels are blocking, since they wait to receive a value
	var ch chan int
	ch = make(chan int)
	go sendNumber(ch, 10)
	fmt.Println(<-ch)

	//close(ch) // Closing channel, means that we can proceed with other computation

	/*
		Getting a value from the channel generates two values:
		value, ok := <- ch
		--value refers to the value that channel returns
		--ok is a boolean value that returns true if the value is generated by a write (e.g. ch <- 10) or false if this value
		is generated by a close (e.g. close(ch)), both examples are explained above
	*/

	go sendNumber(ch, 20)
	value, ok := <-ch

	fmt.Println(value, ok)
	fmt.Println("Done...")

	close(ch)
	//value, ok := <-ch  // If we were to do this, remember that  value is 0 and ok would be false

	ch2 := make(chan int)

	go func(a, b int) {
		c := a + b
		ch2 <- c
	}(1, 2)

	fmt.Println("Computed value", <-ch2)
}

func sendNumber(c chan int, n int) {
	c <- n
}
