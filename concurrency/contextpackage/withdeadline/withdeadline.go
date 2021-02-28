package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	string
}

const duration = 980 * time.Millisecond

func main() {

	dl := time.Now().Add(100 * time.Millisecond)                  // Creating a deadline of 100 milliseconds
	ctx, cancel := context.WithDeadline(context.Background(), dl) // Creating a context withDeadline and a cancel function

	// Calling cancel function is a good practice to avoid keeping the parent context alive longer than necessary
	defer cancel()

	compute := func() <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)
			deadline, ok := ctx.Deadline() // ctx.Deadline checks whether there is a deadline or not, returns false if not
			fmt.Println(deadline.Second())
			if ok {
				if deadline.Sub(time.Now().Add(80*time.Millisecond)) < 0 { // Subtracts an amount of time from the deadline
					fmt.Println("Not sufficient time given... terminating")
				}
			}
			// Simulate work
			time.Sleep(50 * time.Millisecond)

			// Report result
			select {
			case ch <- data{"123"}:
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			}
		}()
		return ch
	}

	ch := compute()
	d, ok := <-ch
	if ok {
		fmt.Println("Work is completed", d)
	}

	ds := time.Now().Add(duration)
	ctx2, cancel2 := context.WithDeadline(context.Background(), ds)
	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel2()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Overslept")
	case <-ctx2.Done():
		fmt.Println(ctx2.Err())
	}
}
