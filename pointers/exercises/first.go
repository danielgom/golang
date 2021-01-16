package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var c *computer

	// compare the pointer variable to a nil value, and say it's nil
	if c == nil {
		fmt.Println("This pointer is nil")
	}

	// create an apple computer by putting its address to a pointer variable

	apple := &computer{brand: "apple"}

	// put the apple into a new pointer variable

	newApple := apple
	// compare the apples: if they are equal say so and print their addresses
	if apple == newApple {
		fmt.Printf("Address of apple is %p, address of newApple is: %p\n", apple, newApple)
	}

	// create a sony computer by putting its address to a new pointer variable

	sony := &computer{brand: "sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses

	if apple != sony {
		fmt.Printf("Address of apple is %p, address of sony is: %p\n",
			apple, sony)
	}

	// put apple's value into a new ordinary variable

	appleVal := *apple

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well

	fmt.Printf("Apple address (pointer address) %p, Address it points to %p\n", &apple, apple)
	fmt.Printf("AppleVal address %p\n", &appleVal)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	// print the values:
	// the value that is pointed by the apple and the new variable

	if *apple == appleVal {
		// if they are equal say so
		fmt.Println("apple and appleVal are equal")

		// print the values:
		// the value that is pointed by the apple and the new variable
		fmt.Printf("apple                     : %+v — appleVal: %+v\n",
			*apple, appleVal)
	}

	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func — print sony's brand
	change(sony, "hp")

	// write a func that returns the value that is pointed by the given *computer
	// print the returned value
	fmt.Printf("appleVal                  : %+v\n", valueOf(apple))

	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("dell's address            : %p\n", newComputer("dell"))
	fmt.Printf("lenovo's address          : %p\n", newComputer("lenovo"))
	fmt.Printf("acer's address            : %p\n", newComputer("acer"))
	
}

func change(c *computer, brand string) {
	c.brand = brand
}

func valueOf(c *computer) computer { // Accepts a pointer value (address value) and returns its pointed value
	return *c
}

func newComputer(brand string) *computer {
	return &computer{brand: brand}
}