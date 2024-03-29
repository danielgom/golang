package main

import (
	"fmt"
	"time"
)

func main() {

	for {

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		alarmed := sec%10 == 0

		for line := range clock[0] {
			if alarmed {
				clock = alarm
			}
			for index, digit := range clock {
				next := clock[index][line]
				if (digit == colon || digit == dot) && sec%2 == 0 {
					next = "     "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		const splitSecond = time.Second / 10
		time.Sleep(time.Second)
		fmt.Print("\033[2J")

	}



}
