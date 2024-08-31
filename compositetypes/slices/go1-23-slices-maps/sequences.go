package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

func main() {

	mySlice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// First part is the same as the second part
	// First part is the compiler result
	backward := slices.Backward(mySlice)
	backward(func(i int, v int) bool {
		fmt.Println(i, v)
		if i == 5 {
			return false
		}
		return true
	})

	// This is the second part, in runtime the following code will be converted
	// to the part 1
	fmt.Println()
	for i, v := range backward {
		if i == 4 {
			break
		}
		fmt.Println(i, v)
	}

	fmt.Println()
	myMap := map[int]int{1: 10, 2: 20, 3: 30, 4: 40, 5: 50, 6: 60, 7: 70, 8: 80, 9: 90, 10: 100}
	for k := range maps.Keys(myMap) {
		fmt.Println(k)
	}

	fmt.Println()
	for v := range maps.Values(myMap) {
		fmt.Println(v)
	}

	firstMap := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"}
	secondMap := map[int]string{6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten"}

	maps.Insert(firstMap, maps.All(secondMap))

	fmt.Println(firstMap)

	firstMapEx2 := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"}
	secondMapEx2 := map[int]string{6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten"}

	InsertToMap(firstMapEx2, secondMapEx2)

	fmt.Println(firstMapEx2)

	oneSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	collectedSlice := slices.Collect(Filter(func(v int) bool {
		return v%2 == 0
	}, slices.Values(oneSlice)))

	fmt.Println(collectedSlice)

}

func InsertToMap[Map ~map[K]V, K comparable, V any](m Map, m2 Map) {
	for k, v := range m2 {
		m[k] = v
	}
}

// Filter returns a sequence that contains the elements
// of s for which f returns true.
func Filter[V any](f func(V) bool, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range s {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}
