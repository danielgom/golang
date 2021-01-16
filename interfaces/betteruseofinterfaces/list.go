package main

import (
	"sort"
	"strings"
)

// Using pointers since method receivers in product are pointers
// This is a recreation of the list in the main interface but without using interfaces, more concse
type list []*product

type byRelease struct {
	list
}


func (l list) String() string { // list is []Printer and each printer contains print(), this should be renamed to print, but for
	// learning purposes we leave it as printItems
	if len(l) == 0 {
		return "There are no elements on the list"
	}

	var str strings.Builder
	for i := range l {
		str.WriteString("* ")
		str.WriteString(l[i].String()) // We need to call String here since we are not using fmt
		str.WriteRune('\n')
		//fmt.Printf("* %s\n", l[i]) // PrintF knows that product has a String method (like a toString() in java), hence
		// printf calls it, THE %S calls the l[i] which is the product
		//l[i].print() // Calls the print function of every product
	}
	return str.String()
}

func (l list) discount(ratio float64) {
	for _, it := range l {
		it.discount(ratio)
	}
}

// TO IMPLEMENT THE SORTER INTERFACE THAT GOLANG HAS WE NEED TO IMPLEMENT THIS METHODS IN THE TYPE WE WANT,
// IN THIS CASE LIST IS GOING TO HAVE SORT, LIST NOW SATISFIES THE SORT INTERFACE

func (l list) Len() int {
	return len(l)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (l list) Less(i, j int) bool {
	return l[i].Title < l[j].Title // What are we comparing? titles
}

// Swap swaps the elements with indexes i and j.
func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i] // How are we going to swap?, this is the answer
}


// The same happens here as it happens with the reverse, we firstly go to this method Less and if necessary we call the
// Less in the list
func (b byRelease) Less(i, j int) bool {
	return b.list[i].Released.Before(b.list[j].Released.Time)
}


// List already has Len, Less and Swap so its fine for the sort.Interface
func byReleaseDate(l list) sort.Interface {
	return &byRelease{l}
}