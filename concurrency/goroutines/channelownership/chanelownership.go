package main

import "fmt"

func main() {

	/*
		Not commonly used, however, this example shows an owner , which means that its responsibilities are creating
		the channel, "populating the channel", and closing the channel
	*/

	owner := func() <-chan int { // This is going to return a send only channel
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		return ch
	}

	consumer := func(ch <-chan int) { // consumer receives a receive only channel in the parameters
		for v := range ch {
			fmt.Println("Received,", v)
		}
		fmt.Println("Done")
	}
	ch := owner()
	consumer(ch)

	// Trying to do a filter of int

	ns := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	filtered := Filter(ns, func(b int) bool {
		return b%2 == 0
	})

	filter := func(a int) bool {
		return a%2 == 0
	}

	filtered2 := Filter(ns, filter)

	fmt.Println(filtered)
	fmt.Println(filtered2)

}

func Filter(a []int, test func(b int) bool) []int {
	fs := make([]int, 0)
	for _, v := range a {
		if test(v) {
			fs = append(fs, v)
		}
	}
	return fs
}
