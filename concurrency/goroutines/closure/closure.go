package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Println("Value of i,", i)
		}()

		fmt.Println("return from function")
		return
	}

	incr(&wg) //Always send a waitGroup as a pointer
	wg.Wait()
	fmt.Println("Done...")

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("The value of i is:", i) // This may not print it on order since the output is printed async
			// so we may see different answers each time we execute the program
		}(i) // Remember to send the value in a for to the anonymous function
	}

	wg.Wait()

	v := fibonacci(3)
	fmt.Println(v)

}

func fibonacci(n int) int {
	current, prev := 0, 1
	for i := 0; i < n; i++ {
		current, prev = current+prev, current
	}
	return current
}
