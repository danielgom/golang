package main

import (
	"fmt"
)

func Print[T any](s []T) {
	for _, val := range s {
		fmt.Println(val)
	}
}

func Map[K, V any](items []K, mapFunc func(K) V) []V {
	result := make([]V, len(items))
	for idx, item := range items {
		result[idx] = mapFunc(item)
	}
	return result
}

func concat[T any](c1, c2 <-chan T) <-chan T {
	r := make(chan T)
	go func(c1, c2 <-chan T, r chan<- T) {
		defer close(r)
		for v := range c1 {
			r <- v
		}

		for v := range c2 {
			r <- v
		}
	}(c1, c2, r)
	return r
}

func main() {
	Print(Map([]int{1, 2, 3, 4, 5}, func(item int) bool {
		return item%2 == 0
	}))

	c1 := make(chan string, 2)
	c2 := make(chan string, 2)
	go func() {
		c1 <- "hello"
		c1 <- ", "
		close(c1)

		c2 <- "world"
		c2 <- "!"
		close(c2)

	}()

	c3 := concat(c1, c2)
	for elem := range c3 {
		fmt.Print(elem)
	}
}
