package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/*
		Pipelines are used to process streams or batches of data, it also helps to the separation of concerns
		this being said, each pipeline is composed by many stages and each stage is represented by a go routine
		if a stage is big enough then we can scale the go routines up in that stage, in order to avoid affecting
		the whole pipeline
	*/

	// TODO generator() -> square() -> print
	rand.Seed(time.Now().Unix())
	ch := generator(rand.Perm(100)...)

	out := square(ch)

	for n := range out {
		fmt.Printf("%d, ", n)
	}

	// TODO Better way to call the pipeline

	for n := range square(generator([]int{1, 2, 3, 4, 5}...)) {
		fmt.Println(n)
	}

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
