package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	phones := map[string]string{

		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}

	products := map[int]bool{
		617841573: true,
		879401371: false,
		576872813: true,
	}

	multiPhones := map[string][]string{
		"bowen": {"202-555-0179"},
		"dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": {"03489940240", "03489900120"},
	}

	basket := map[int]map[int]int{
		100: {617841573: 4, 576872813: 2},
		101: {576872813: 5, 657473833: 20},
		102: {},
	}

	fmt.Printf("%#v\n", phones)
	fmt.Printf("%#v\n", products)
	fmt.Printf("%#v\n", multiPhones)
	fmt.Printf("%#v\n", basket)

	fmt.Printf("%T\n", phones["dulin"]) // slices
	fmt.Println(phones["dulin"])

	if !products[617841573] {
		fmt.Println("Value exists")
	} else {
		fmt.Println("It does not exists")
	}

	fmt.Println(multiPhones["greco"][1])
	fmt.Println(basket[101][576872813])

	/*
		basket:

			1:
				1: 1
				2: 2
			2:
				1: 1
				2: 2

			3:
				1: 1
				2: 2
				3: 5
	*/

	houses := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory", "bobo"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "bobo", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	delete(houses, "bobo")

	fmt.Println(houses)

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}

	house, students := args[0], houses[args[0]]
	if students == nil { // Same way to check for a value as the other one, here we check if value is null
		fmt.Printf("Sorry. I don't know anything about %q.\n", house)
	}

	if _, exists := houses[args[0]]; !exists { // here we just add the exists and check if it exists
		fmt.Printf("Sorry. I don't know anything about %q.\n", house)
		return
	}

	// only sort the clone
	clone := append(make([]string, 0, len(students)), students...)
	sort.Strings(clone)

	fmt.Printf("~~~ %s students ~~~\n\n", house)
	for _, student := range clone {
		fmt.Printf("\t+ %s\n", student)
	}

}
