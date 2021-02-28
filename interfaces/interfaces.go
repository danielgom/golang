package main

import (
	"fmt"
)

func main() {

	var (
		moby    = book{product{title: "", price: 19.99}, 118281600}
		diablo4 = game{product{title: "Diablo 4", price: 99.99}}
		cyberCG = game{product{title: "Cybeer cg", price: 199.99}}
		puzz    = puzzle{product{title: "Interesting", price: 12.49}}
	)

	// Remember that  type list []Printer (slice of printers, printers is an interface)
	var store list
	store = append(store, &moby, &diablo4, &cyberCG, &puzz) // Why do we use the addresses instead of the values?
	/* Answer: since the print() method in every type has a pointer receiver *type, therefore, we need to use an address
	Explanation, basically we have a []item, item interface has a print() method in it, just like Java language
	we need to implement that method, hence , book, game and puzzle types have that method implemented, if they didnt
	an error would happen which warns us about a type not implementing the methods in the item interface

	Remember: A type satisfies an interface automatically when it has all the methods of the interface without explicitly
	specifying it, we do not need to use implements like java, in this case , book, game and puzzle have the method print
	therefore these are implicitly item types, that's why we can append it to the list which is []Item
	*/
	store.printItems()

	fmt.Println()
	Print(store)

	// Notice that we can compare structs of the same type
	fmt.Println(diablo4 == game{product{title: "Diablo 4", price: 99.99}}) // True
	fmt.Println(diablo4 == cyberCG)                                        // False
	//fmt.Println(diablo4 == puzz) Notice that we cannot compare structs of different type

	// We can compare interaces in this case
	// Notice in the following code we are comparing a book struct type with a game struct type
	fmt.Println(store[1] == store[2]) // It works and result is false
	fmt.Println(store[1] == &diablo4) // result is true

	fmt.Println()
	fmt.Println()
	fmt.Println()

	var i item

	// Zero value of an interface is nil
	fmt.Println(i)

	i = store[0] // Getting the first value of the []item list , it returns an item instead of the moby book
	i.print()    // We can just call the print method here since it is type item not type book

	i = &diablo4
	i.print()

	// Return/extract the dynamic value

	var newBook = store[0].(*book) // This is the way to cast the from item to book
	newBook.print()
	newBook.changeTitle("New moby dick")
	newBook.print()

	//var newGame = store[1].(*book) // Throws an error because store[1] contains an item which is from a type game not book
	var newGame = store[1].(*game) // Since it is a game then we need to convert back to a game type
	fmt.Println()
	newGame.print()
	newGame.discount(0.30)
	newGame.print()

	fmt.Println()
	// If we want to apply discounts only on the games of the []item list, check the bounded discount method
	store.printItems()
	store.discount(.40)
	fmt.Println()
	store.printItems() // As we can see methods got a discount whereas any other type has the same value
	store.discountPerf(.10)
	store.printItems()

	fmt.Println()

	// DO NOT USE EMPTY INTERFACES UNLESS REALLY NECESSARY , USING IT MAY LEAD IN UNMAINTAINABLE CODE

	// We create another item list remember that list is []item, which only needs a print() method to be implemented
	// Each book has different values of interfaces

	// Since we embedded the product type then we need to specify it on every creation of game, book , puzzle and toy
	// We need the & since product uses pointers on the receivers

	store2 := list{
		&book{product: product{title: "mobydick", price: 19.99}, published: "733622400"},      // Published is an empty interface type interface{}
		&book{product: product{title: "the last breath", price: 29.99}, published: 118281600}, // therefore we can have anything there like ints, strings and booleans
		&book{product: product{title: "It", price: 9.99}, published: false},                   // or any other value
		&book{product: product{title: "It", price: 9.99}, published: nil},                     // it can be nil as well
		&game{product: product{title: "Diablo 4", price: 99.99}},
		&game{product: product{title: "Cybeer cg", price: 199.99}},
		&puzzle{product{title: "Interesting", price: 12.49}},
	}
	store2.printItems()

	// Type switch
	fmt.Println()

	// Even if the toy does not have the print and method, it takes these methods from the embedded product struct type
	t := &toy{product{title: "yoda", price: 150.00}}
	t.discount(.10)
	t.print()

	b := book{product: product{title: "New book", price: 19.99}, published: "733622400"}
	b.print()

	// In the end we did not have to change anything in the item type interface since these types, book, game, toy and puzzle
	// have the product type, and this type implicitly implements the item interface with the print and the discount methods

	fmt.Println()
	fmt.Println()
	fmt.Println()

	var newItem item
	newItem = &book{product: product{title: "It", price: 9.99}, published: false}
	newItem.print()
	b2 := newItem.(*book)
	b2.discount(.20)
	b2.print()
}
