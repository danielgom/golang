package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(2 * time.Second)

	<-timer.C

	fmt.Println("Timer 1 fired")
	timer2 := time.NewTimer(time.Second / 2)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	time.Sleep(time.Millisecond * 300)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
