package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	// Not a good way to create a map
	//var dict map[string]string // Not initialized, a non initialized map cannot add keys and values but we can
	//check its length which is going to be 0

	// Need to initialize the map

	args := os.Args[1:]

	dict := map[string]string{} // The{} initializes the map

	dict["hello"] = "nope" // Nil map error

	fmt.Printf("%#v\n", dict)
	fmt.Println(dict[args[0]])

	// Another way to create a map

	engTranslate := map[string]string{
		"good":      "bueno",
		"bad":       "malo",
		"excellent": "excelente",
	}

	fmt.Println(engTranslate)

	// If we try too look for a key that isn't in a map, then the response is the zero value of the key
	// for example:

	intMap := map[int64]string{
		1: "supp",
		2: "where",
		3: "yep",
	}

	fmt.Println(intMap[1]) // returns supp
	fmt.Println(intMap[4]) // returns the 0 value onf an int64 which is 0

	// if we want to know if something is in a map this is the best way to do it

	val, exists := intMap[1] // Val returns the val if it exists, if it does not exist , returns the 0 value of the key type of the map
	// exists return whether the key we want to retrieve exist or not (true or false)
	fmt.Println(val, exists) // prints "supp" and true
	val, exists = intMap[4]
	fmt.Println(val, exists) // prints "" and false

	//Normal way to do it

	if val, exists := intMap[10]; exists {
		fmt.Println("value: ", val, " exists")
	} else {
		fmt.Println("It does not exists")
	}

	// Comparing maps

	firstMap := map[int]string{

		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
	}

	secondMap := map[int64]string{

		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
		206: "Sumit",
	}

	thirdMap := map[int]string{

		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
	}

	fmt.Println(firstMap)
	fmt.Println(secondMap)
	fmt.Println(thirdMap)

	fmt.Println("Is map 1 equal to map 2?: ", reflect.DeepEqual(firstMap, secondMap))
	fmt.Println("Is map 1 equal to map 3?: ", reflect.DeepEqual(firstMap, thirdMap))
	fmt.Println("Is map 2 equal to map 3?: ", reflect.DeepEqual(secondMap, thirdMap))

	// Map works the same way as slices, slices contain the pointer to the backing array
	// Map contains the pointer to the actual values of the map

	thisMap := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	nextMap := thisMap //This makes that nextMap points to the same values as thisMap (this is super efficient just like slices)
	// but chaning any vaue in nextMap, affects thisMap value

	nextMap["fourth"] = 4
	fmt.Println(thisMap) // map[first:1 fourth:4 second:2 third:3]
	fmt.Println(nextMap) // map[first:1 fourth:4 second:2 third:3]

	// The best way to do it is to create it with the make command

	newMap := make(map[string]int, len(thisMap)) // This is different from a slice creation, it does not necessary mean that length is
	// the length value we have given, it means that this map is going to contain AT LEAST , the length value

	fmt.Println(len(newMap)) // Prints 0

	// REMEMBER THAT IF A MAP ALREADY CONTAINS A KEY AND WE TRY TO ADD THAT SAME KEY, THE VALUE IS GOING TO BE UPDATED

	// How to delete a map
	nextMap = nil        // This just deletes the pointer to the map, however the map still exists
	fmt.Println(thisMap) // still contains map[first:1 fourth:4 second:2 third:3]
	fmt.Println(nextMap) // map[] since we deleted the pointer

	// Deleting the whole map

	// **********************************************************
	for x := range thisMap {
		delete(thisMap, x) // deletes a key from a map
	}

	// **********************************************************

	/*
		The for above deletes every key from the map however it does not loop over the whole map, golang makes an efficient
		clear of the map calling a mapclear() function behind de scene, this is very efficient, so do not hesitate to delete
		a map this way
	*/

	fmt.Println(thisMap) // map[] deleted the whole map

}
