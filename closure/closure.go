package main

import "fmt"

func main() {

	mySuperFunction := func(num int) func() int {
		fmt.Println("Doing something before executing next function")
		fmt.Println("We are going to sum til", num)

		return func() int {
			var sum int
			for range num {
				sum++
			}
			fmt.Println("We are done summing", sum)
			return sum
		}
	}

	defer mySuperFunction(2)()
}
