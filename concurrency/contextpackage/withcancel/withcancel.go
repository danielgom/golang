package main

import (
	"context"
	"fmt"
)

func main() {

	/*
		Created a cancellable context, this cancellable context returns the context and a cancellable function whenever cancel
		function is called we send a signal to the context and that signal is going to trigger the done function
	*/

	ctx, cancel := context.WithCancel(context.Background())

	for n := range generator(ctx) {
		fmt.Println(n)
		if n == 100 {
			cancel() // Send the cancel signal
		}
	}

}

func generator(ctx context.Context) <-chan int {
	ch := make(chan int)
	n := 1
	go func() {
		defer close(ch)
		for {
			select {
			case ch <- n:
				n++
			case <-ctx.Done(): // When we receive a cancel signal from ctx, done is raised
				return // Return when we receive the cancel signal
			}
		}
	}()
	return ch
}
