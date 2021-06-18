package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

func main() {

	args := os.Args[1:]

	am, err := strconv.Atoi(args[0])
	am += 1

	if err != nil {
		log.Fatal(err)
	}

	rands := rand.Perm(am)
	c := make(chan int64, 5)

	start := time.Now()

	if v := am % 5; v != 0 {
		r := len(rands) - v
		for x := 1; x <= 5; x++ {
			if x == 5 {
				go sum(rands[r/5*(x-1):r/5*x+v], c)
			} else {
				go sum(rands[r/5*(x-1):r/5*x], c)
			}
		}

	} else {
		for x := 1; x <= 5; x++ {
			go sum(rands[len(rands)/5*(x-1):(len(rands)/5)*x], c)
		}
	}
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(duration.Seconds())

	// Retrieve


	var result int64

	for x := 0; x < 5; x++ {
		atomic.AddInt64(&result, <-c)
	}
	fmt.Println(result)
	/*
	ts := 0
	for i := 0; i < cap(c); i++ {
		n := <-c
		fmt.Println(n)
		ts += int(n)
	}
	fmt.Println(ts)

	 */

}

func sum(sl []int, c chan<- int64) {
	s := 0
	for i := range sl {
		s += sl[i]
	}
	c <- int64(s)
}
