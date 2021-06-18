package main

import (
	"fmt"
	"sync"
)

func main() {

	/*
		Data race is whenever two or more go routines are trying to modify a value of a variable leading into misleading
			results
	*/

	/*
		Following example shows a variable n which starting value is 0, we are adding +1 in the go routine and in the
		main go routine, we can have multiple results of this, we would like to have a result of 2 however, that does
		not happen frequently
	*/
	n := 0

	go func() {
		n++
	}()
	n++
	fmt.Println(n)

	/*
		Fixing the issue
	*/

	n2 := 0
	c := make(chan int)
	go func() {
		n2 := 0
		n2++
		c <- n2
	}()
	n2 = <-c
	n2++
	fmt.Println(n2)

	/*
		Fixing the issue with waitGroups and mutex
	*/

	n3 := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		add(&n3, &mu)
	}()
	add(&n3, &mu)
	add(&n3, &mu)
	add(&n3, &mu)

	wg.Wait()
	fmt.Println(n3)

	/*
		Fixing the issue with only mutex
	*/

	var mutex sync.Mutex

	try := 0
	go add(&try, &mutex)
	go add(&try, &mutex)
	go add(&try, &mutex)
	go add(&try, &mutex)

	fmt.Println(try)

}

func add(a *int, mu *sync.Mutex) {
	mu.Lock()
	*a++
	mu.Unlock()

}
