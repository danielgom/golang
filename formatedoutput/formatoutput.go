package main

import "fmt"

func main() {

	fmt.Println("Hello \\n\"how are you?\" ")

	var brand = "google"

	fmt.Printf("%s", brand)
	fmt.Printf("%q\n", brand)

	// using integers

	speed := -22
	fmt.Printf("This is the speed: %d\n", speed)

	var (
		howManyHours = 23
		isItReady    = true
		piValue      = 3.141594
		comment      = "No comments at all"
	)

	fmt.Printf("There are %s, being the fact that Ready is %t and counting some pi values with this value %g, in some %d hours\n",
		comment, isItReady, piValue, howManyHours)

	// precision

	fmt.Printf("Pi value: %.0f\n", piValue)
	fmt.Printf("Pi value: %.1f\n", piValue)
	fmt.Printf("Pi value: %.2f\n", piValue)
	fmt.Printf("Pi value: %.3f\n", piValue)
	fmt.Printf("Pi value: %.4f\n", piValue)
	fmt.Printf("Pi value: %.5f\n", piValue)
	fmt.Printf("Pi value: %.6f\n", piValue)
	fmt.Printf("Pi value: %.7f\n", piValue)

	// types
	fmt.Printf("Type of hours %T\n", howManyHours)
	fmt.Printf("Type of is it ready %T\n", isItReady)
	fmt.Printf("Type of pivalue %T\n", piValue)
	fmt.Printf("Type of hours %T\n", comment)

}
