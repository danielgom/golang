package main

import (
	"fmt"
	"sync"
)

func producer(id int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range 3 {
			ch <- id*10 + i
		}
	}()
	return ch
}

func fanIn(inputs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	output := make(chan int)

	cp := func(c <-chan int) {
		defer wg.Done()
		for val := range c {
			output <- val
		}
	}
	
	wg.Add(len(inputs))
	for _, ch := range inputs {
		go cp(ch)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func main() {
	// Step-1 Producer
	ch1 := producer(1)
	ch2 := producer(20)

	// Step-1 FanIn
	mergedOutputChannel := fanIn(ch1, ch2)

	// Step-3 Consumer
	for val := range mergedOutputChannel {
		fmt.Println(val)
	}
}
