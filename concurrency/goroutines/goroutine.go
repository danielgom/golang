package main

import (
	"fmt"
	"time"
)

func main() {

	fun("direct call")

	// go routine call
	go fun("go routine-1")

	// go routine anonymous function/function literal
	go func() {
		fun("go routine-2")
	}()

	// go routine function/value call
	fv := fun
	go fv("go routine-3")

	fmt.Println("wait for go routines")
	time.Sleep(5 * time.Millisecond)
	fmt.Println("done...")

}

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}
