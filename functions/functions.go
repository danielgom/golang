package main

import (
	"fmt"
	"math"
)

type person struct {
	name    string
	numbers []int
}

type professional struct {
	name string
	jobs map[int]string
}

func main() {

	sqrt := math.Sqrt(22)
	fmt.Println(sqrt)

	daniel := person{
		name:    "Daniel Gomez Alvarez",
		numbers: []int{1, 2, 3, 4, 5},
	}
	fmt.Println("Before func (slice): ", daniel)
	changeStructWithSlice(daniel)
	fmt.Println("After func (slice): ", daniel)

	daniel2 := daniel
	daniel2.numbers = append(daniel2.numbers, 7, 8, 9, 10) // Creates a new backing array and that's why original is not changed
	fmt.Println("After append change to the copy (slice)", daniel)
	pro := professional{
		name: "Daniel Gomez Alvarez",
		jobs: map[int]string{1: "TCS", 2: "IBM", 3: "Google"},
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println("Before func (map): ", pro)
	changeStructWithMap(pro)
	fmt.Println("After func (map): ", pro)

	pro2 := pro
	pro2.jobs[5] = "Who knows"
	fmt.Println("After add change to the copy (map)", pro)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	aldo := person{
		name:    "Aldo Coll Heredia",
		numbers: []int{1, 2, 3, 4, 5},
	}
	fmt.Println("Testing with a function which returns the same struct with changes (slice check)")
	fmt.Println("Before func (slice): ", aldo)
	aldo = changeStructWithSliceReturn(aldo)
	fmt.Println("After func (slice)", aldo)

}

func changeStructWithSlice(p person) { // Whenever we want to change a slice with full capacity we need to return it
	p.numbers = append(p.numbers, []int{7, 8, 9, 10}...)
	fmt.Println("This is the copy in the func (slice): ", p)
}

func changeStructWithMap(pro professional) {
	pro.jobs[4] = "This is a mystery"
	fmt.Println("This is the copy in the func (map): ", pro)
}

func changeStructWithSliceReturn(p person) person {
	p.numbers = append(p.numbers, 7, 8, 9, 10)
	fmt.Println("This is the copy in the func (slice): ", p)
	return p
}
