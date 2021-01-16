package main

import "fmt"

type book struct {
	title string
	price money
}

// USE POINTER RECEIVERS WHEN WE ARE GOING TO CHANGE THE STATE , OR WHEN DATA IS TOO LARGE

func (b *book) printBook() { // Do no use methods this way, methods are meant to be used and change the state of the object
	//Only mistake here is not changing the state of the object
	fmt.Printf("Title: %s, Price: %s\n", b.title, b.price.string())
}

func (b book) printBookValue() { // Remember, only use methods to change state, not to just print the struct
	fmt.Printf("Title: %s, Price: %s\n", b.title, b.price.string())
}

func (b *book) changeTitle(t string) { // If we do not add the * then we would need to return the complete b struct and save it
	b.title = t
}

func (b book) changeTitleValue(t string) book { // Need to return the book and reassign since we are not using pointers
	b.title = t
	return b
}


