package main

import (
	"fmt"
	"github.com/dgomez/learngo/first/explain/compositetypes/slices/slicespractice/api"
	"io/ioutil"
	"math/rand"
)

var temps = []int{5, 10, 3, 25, 45, 80, 90}

func main() {

	received := Read(0, 3)
	fmt.Println(received)
	fmt.Println(len(received))
	received = append(received, []int{1, 3}...)

	fmt.Println(All())
	fmt.Println(received)
	fmt.Println(cap(received))

	randoms := rand.Perm(500)
	fmt.Println(randoms)

	api.Report()

	// returns a slice with 10 million elements.
	// it allocates 65 MB of memory space.
	millions := api.ReadMillions()

	// -----------------------------------------------------
	// ✪ ONLY CHANGE THE CODE IN THIS AREA ✪

	last10 := make([]int, 0, len(millions[len(millions) - 10:]))

	last10 = append(last10, millions[len(millions)-10:]...)

	millions = last10

	fmt.Printf("\nLast 10 elements: %d\n\n", last10)

	// ✪ ONLY CHANGE THE CODE IN THIS AREA ✪
	// -----------------------------------------------------

	api.Report()

	// don't worry about this code.
	fmt.Fprintln(ioutil.Discard, millions[0])

}

func Read(start, stop int) []int {
	//
	// The third index prevents the `main()` from
	// overwriting the original temps slice's
	// backing array. It limits the capacity of the
	// returned slice. See the full slice expressions
	// lecture for more details.
	//
	portion := temps[start:stop:stop]

	return portion
}

// All returns the temps slice
func All() []int {
	return temps
}
