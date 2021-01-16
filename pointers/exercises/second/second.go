package main

import "fmt"

func main() {

	// EXERCISE: Swap
	//
	//  Using funcs, swap values through pointers, and swap
	//  the addresses of the pointers.
	//
	//
	//  1- Swap the values through a func
	//
	//     a- Declare two float variables
	a, b := 2.25 , 3.45
	fmt.Printf("a : %g   b : %g\n", a, b)
	//
	//     b- Declare a func that can swap the variables' values
	//        through their memory addresses
	//
	//     c- Pass the variables' addresses to the func
	//
	//     d- Print the variables
	//
	swap(&a, &b)
	fmt.Println("After function")
	fmt.Printf("a : %g   b : %g\n", a, b)
	//
	//  2- Swap the addresses of the float pointers through a func
	//
	//     a- Declare two float pointer variables and,
	//        assign them the addresses of float variables
	pa, pb := &a, &b
	fmt.Printf("pa: %p   pb: %p\n", pa, pb)
	fmt.Printf("pa: %g   pb: %g\n", *pa, *pb)
	pa, pb = swapAddress(pa, pb)
	fmt.Printf("pa: %p   pb: %p\n", pa, pb)
	fmt.Printf("pa: %g   pb: %g\n", *pa, *pb)
	//
	//     b- Declare a func that can swap the addresses
	//        of two float pointers
	//
	//     c- Pass the pointer variables to the func
	//
	//     d- Print the addresses, and values that are
	//        pointed by the pointer variables
	//
	// ---------------------------------------------------------
}

func swap(first, second *float64) { //Passing pointer value
	*first, *second = *second, *first //Remember that since we are passing we need to use * to retrieve its value
}

func swapAddress(first, second *float64) (*float64, *float64) { //Must return if we want to swap addresses since we create other values inside here
	//first, second = second, first
	return second, first
}
