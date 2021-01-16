package main

import "fmt"

type computer struct {
	brand *string
}

func main() {
	// ---------------------------------------------------------
	// EXERCISE: Fix the crash
	//
	// EXPECTED OUTPUT
	//
	//   brand: apple
	// ---------------------------------------------------------

	//var c *computer
	c := &computer{}
	change(c, "apple")
	fmt.Printf("brand: %s\n", *c.brand) // print *c.brand which contains "apple"


}

func change(c *computer, brand string) {
	c.brand = &brand //Change value of c.brand to point to brand string address
}
