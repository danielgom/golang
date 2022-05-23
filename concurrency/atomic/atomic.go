package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	/*
		Atomic is used to basically make concurrent operations safe
	*/

	runtime.GOMAXPROCS(10) // We are telling golang to use max 4 cpus

	var counter uint64
	var wg sync.WaitGroup

	for x := 0; x < 50; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 100; c++ {
				//counter++ Notice that we are not simply adding one to the counter
				atomic.AddUint64(&counter, 1) // We are passing the value as a reference and telling to add 1 to the unit
				// Atomic provides a safer way to do concurrent operations
				// After adding the atomic operation, we get a consistent result and the result we previously expected
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	/*
		This past example shows interesting results, we are declaring many go routines to modify our counter variable,
		hence the result of these operations may be corrupted, that's why we need a safe way to run these operations concurrently
		NOTE: We can use Mutex however, we would not be concurrent, using atomic can change our result
	*/
}
