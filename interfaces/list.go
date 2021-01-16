package main

import (
	"fmt"
)

//type list []*game

type list []item

// We want to make a type that can contain more than just []*game game addresses, therefore, we created an interface
type item interface {
	print()
	discount(ratio float64)
}

func (l list) printItems() { // list is []Printer and each printer contains print(), this should be renamed to print, but for
	// learning purposes we leave it as printItems
	if len(l) == 0 {
		fmt.Println("There are no elements on the list")
		return
	}
	for i := range l {
		l[i].print() // This calls the print method of printer interface and each implementation
	}
}

func Print(li []item) { // This is not the best way to do it, better use the method instead of the function
	if len(li) == 0 {
		fmt.Println("There are no elements on the list")
		return
	}
	for i := range li {
		li[i].print() // This calls the print method of printer interface and each implementation
	}
}

// If we want to use for example the discount method that comes in the game types we can do the following
// However there is a more efficient wat to do so, check discountPerf method
func (l list) discount(ratio float64) {
	for i := range l {
		g, isGame := l[i].(*game) // Here what returns is false for every type that is not game time in the list of items
		if !isGame {
			continue // Jumps the following call of discount for non-game types
		}
		g.discount(ratio)
	}
}

// discountPerf checks for behaviour, in this case is checking if any item in the []item list has a discount method
// if this method is encountered then it calls it, only the game type has this method
func (l list) discountPerf(ratio float64) {
	/*
		type discounter interface { // We create a local interface that which needs the discount method to be satisfied
			discount(float64) // We just need to check the discount method with float64 parameter in it
		}
	*/
	for _, it := range l {
		it.discount(ratio)
	}
}
