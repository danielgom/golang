package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		Select let us wait for MULTIPLE channel operations, combining channels with the select statement is common and a
		really powerful feature of golang language
	*/

	// Creation of channels
	c1 := make(chan string)
	c2 := make(chan string)

	// Execution of an operation within channels
	go func() {
		time.Sleep(time.Second * 1) // simulate a ~1 second operation
		c1 <- "Operation one"
	}()

	go func() {
		time.Sleep(time.Second * 2) // simulate a ~2 second operation
		c2 <- "Operation two"
	}()

	/*
		Following code will process the receive on MULTIPLE channels depending on the order the values arrive on the channel
		rather on how they are coded, one example would be, changing the channel one operation to take 3 seconds instead
		of 1 second, this would produce a different result. Operation two would be completed faster
		NOTE: Operations in this example take about 2~ seconds to complete, owing to the fact that, these operations are
		being executed concurrently, therefore, we don't add up operation time
	*/

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received,", msg1)
		case msg2 := <-c2:
			fmt.Println("Received,", msg2)
		}
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
