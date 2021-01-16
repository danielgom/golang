package main

import (
	"fmt"
)

func main() {

	fmt.Println("Pointers")
	var counter int = 100
	P := &counter
	V := *P
	fmt.Println(counter)
	fmt.Println(P)
	fmt.Println("What is inside the P pointer?: ", V)
	fmt.Printf("Pointer type %T\n", P) // *int
	fmt.Printf("Value type %T\n", V)   // int

	// Indirectly change a value through a pointer
	*P = 200 // Since P is pointing to the &counter (address) we can change its value
	// doing a *P assignment to 200
	fmt.Println(counter) // From 100 to 200 indirectly, (memory address is the same)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println("Working pointers")

	var counter2 = 1000
	PO := &counter2
	VO := *PO
	fmt.Printf("Counter2 value %d\n", counter2)                      // int
	fmt.Printf("Counter2 address %v\n", &counter2)                   // int
	fmt.Printf("Pointer PO type %T\n", PO)                           // *int
	fmt.Printf("Pointer PO value (address of saved value) %p\n", PO) // address of the counter 2 value
	fmt.Printf("Pointer PO address (own pointer address) %v\n", &PO) // hex 0x0...
	fmt.Printf("Pointer PO value it is pointing to %d\n", *PO)       // 1000
	fmt.Printf("VO value it is pointing to %d\n", VO)                // int

	fmt.Println()
	fmt.Println()
	fmt.Println()

	/*
		There are two ways of working with functions and return values
	*/

	// Calling the normal function we need to return the value and assign the value to see the changes
	fmt.Println("Working with functions")
	counter2 = passVal(counter2) // we need to assign the returned value
	fmt.Println("Value after the function is called: ", counter2)

	// Calling the pointer function is completely different
	passPointerValue(&counter2)
	fmt.Println("Value after function is called without assignment:", counter2)

}

func passVal(i int) int { // Remember that functions internally create variables and changes them
	i = 2000
	fmt.Println("Value inside the normal function: ", i)
	return i
}
func passPointerValue(pi *int) {
	// This creates a pointer each time it is called (all functions work this way)
	// however every pointer is pointing to the same address that comes from pi
	*pi = 3000
	fmt.Println("Value inside the pointer function", *pi)
}
