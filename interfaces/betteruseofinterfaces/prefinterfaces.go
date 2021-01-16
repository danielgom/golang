package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

// Do not make use of interfaces if you do no truly need them
// Use of single responsibility
func main() {

	// Stringer interface, we just need to have a String method which returns a string to have like a toString() method
	// like java

	// Sorter interface, when trying to sort a list, we need to implement the Len, Swap, and Less functions in order to
	// satisfy the sorter interface

	// Marhsaler interface, when trying to print to JSON there is a way to do it , MarshalJSON() is the method that
	// the Marshaler Interface needs in order to function, when we call MarshalIndent() it looks for for the MarshalJSON() in the type
	// if it has the method implemented then it calls it

	store := list{
		{Title: "mobydick", Price: 19.99, Released: toTimestamp("733622400")},
		{Title: "the last breath", Price: 29.99, Released: toTimestamp(118281600)},
		{Title: "It", Price: 9.99, Released: toTimestamp(false)},
		{Title: "It 2", Price: 9.99, Released: toTimestamp(nil)},
		{Title: "Cybeer cg", Price: 199.99},
		{Title: "Interesting", Price: 12.49},
	}

	//store.print()
	store.discount(0.5)
	fmt.Println()
	//store.print()
	fmt.Println(store) // Calls the String on the store, store is a list type

	// Created a bottle example, with a String method, stringer interface is implicitly implemented and thus whenever
	// we print the example which is a bottle , it prints the type "beautifully"
	example := bottle{capacity: 44.5, height: 1.2}
	fmt.Println(example)

	// Money has a string value, println firstly calls the String then we attach the "I have" to the string we had
	m := money(22)
	fmt.Println(m)
	fmt.Println("I have", m) // Result: I have $22.00

	fmt.Println(store)

	fmt.Println("Store sorted by product name")
	sort.Sort(store)
	fmt.Println(store)

	// To reverse we just need to call the sort.Reverse, basically we firstly call the Less in the reverse interface
	// which consequently calls the Less function in the type we are implementing it (list) but in the reverse order
	// please notice the i and j values reversed

	fmt.Println("Store sorted by product name (reverse)")
	sort.Sort(sort.Reverse(store))
	fmt.Println(store)

	// Creating our own sort

	fmt.Println("Store sorted by release date")
	sort.Sort(byReleaseDate(store))
	fmt.Println(store)

	fmt.Println("Store sorted by release date (reverse)")
	sort.Sort(sort.Reverse(byReleaseDate(store)))
	fmt.Println(store)

	data, err := json.MarshalIndent(store, "", "  ") // Calls the marshal in each type, only timestamp has it implemented
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	const datas = `[
  {
    "Title": "mobydick",
    "Price": 9.995,
    "Released": 733622400
  },
  {
    "Title": "the last breath",
    "Price": 14.995,
    "Released": 118281600
  },
  {
    "Title": "It",
    "Price": 4.995,
    "Released": 0
  },
  {
    "Title": "It 2",
    "Price": 4.995,
    "Released": -62135596800
  },
  {
    "Title": "Interesting",
    "Price": 6.245,
    "Released": -62135596800
  },
  {
    "Title": "Cybeer cg",
    "Price": 99.995,
    "Released": -62135596800
  }
]`
	var s list
	errs := json.Unmarshal([]byte(datas), &s)
	if errs != nil {
		log.Fatal(errs)
	}
	fmt.Println(s)

	fmt.Println(example)

	pnew := product{
		Title:    "nose",
		Price:    20.20,
		Released: toTimestamp(213123123),
	}
	fmt.Println(pnew)

}
