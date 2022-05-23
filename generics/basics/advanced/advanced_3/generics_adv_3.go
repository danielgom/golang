package main

import "fmt"

// type sets
// ~ means that the interface can accept exactly the type or any type that has this type as it's underlined type

type adder interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 |
		~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

func aggregate[T adder](items []T, seed T) T {
	result := seed
	for _, item := range items {
		result += item
	}
	return result
}

// Idx fulfills the adder constraint because it's underlying type uint is part of adder's type list
type idx uint

func main() {

	// Following examples with permissive types for adder

	fmt.Println(aggregate([]int{1, 2, 3, 4}, 32))
	fmt.Println(aggregate([]string{"4", "2"}, " "))
	fmt.Println(aggregate([]idx{21, 23}, 2))

	// The next line will fail because hero cannot satisfy adder,
	//fmt.Println(aggregate([]hero{{name: "superman", canFly: true}}, hero{name: "batman", canFly: false}))

}
