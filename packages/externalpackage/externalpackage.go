package main

import (
	"fmt"
	"github.com/dgomez/master_go/packages/greet"
	"packages/numbers"
)

func main() {

	// Using external package created
	fmt.Println("Is number 22 even?,", numbers.Even(22))
	fmt.Println("Is number 61 prime?", numbers.IsPrime(61))
	greet.Greeter()

}
