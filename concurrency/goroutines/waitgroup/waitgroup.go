package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

	var data int
	var wg sync.WaitGroup

	// Remember to pass wait group value as a pointer

	// Add a new go routine
	wg.Add(1)
	go func() {
		defer wg.Done() // wait for every action to be done before calling done
		data++
	}()

	wg.Wait() // Wait for all go routines to finish in this case only 1
	fmt.Println("The value of the data is:", data)
	fmt.Println("Done...")

	for i := 0; i <= 3; i++ {
		wg.Add(1)         // Add 1 to the wait group meaning that we add 1 each time this for is called
		go worker(i, &wg) // We add one go routine each time this is called , in this case
		// we add 1 to wait group and we add 1 go routine
	}

	// wait for all the counter of wg goes to 0, like before, wait all go  routines to finish
	wg.Wait()

	var answer1 string
	fmt.Printf("Please insert a string: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer1)

}

// Passing the waitGroup value as a pointer allows us to avoid creating a copy of the waitGroup

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Remember that this is run after everything else is run
	fmt.Println("Worker starting,", id)
	time.Sleep(time.Millisecond * 20)
	fmt.Println("Worker done,", id)

}
