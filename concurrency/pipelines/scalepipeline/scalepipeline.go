package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	in := generator(rand.Perm(10)...)
	ch1 := square(in)
	ch2 := square(in)

	for v := range merge(ch1, ch2) {
		fmt.Println("Value", v)
	}

}

func merge(cs ...<-chan int) <-chan int {
	// Merge all channels into a single channel
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for i, _ := range nums {
			out <- nums[i]
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
