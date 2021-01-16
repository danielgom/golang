package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	days   = 1
	months = 3
	years  = 1
)

func main() {

	var books [4]string

	books[0] = "Hello"
	books[1] = "This may be a really big string for only 10 bytes or maybe not"
	books[2] = "nope"
	books[3] = "it might be"

	fmt.Printf("Books type: %T\n", books)
	fmt.Printf("Books size: %d\n", len(books))
	fmt.Println("Books: ", books)
	fmt.Printf("Books type: %q\n", books)
	fmt.Printf("Books type: %#v\n", books)
	fmt.Println("Books: ", &books[0])
	fmt.Println("Books: ", &books[1])
	fmt.Println("Books: ", &books[2])
	fmt.Println("Books: ", &books[3])

	var integers [5]int32

	fmt.Println("integers: ", &integers[0])
	fmt.Println("integers: ", &integers[1])
	fmt.Println("integers: ", &integers[2])
	fmt.Println("integers: ", &integers[3])
	fmt.Println("integers: ", &integers[4])

	fmt.Printf("Combining: %s", books[0]+" once "+books[3]+"\n")

	for _, x := range books { // Creates a copy of the original, not good for iterating over big arrays
		fmt.Println(x)
	}

	for x := range books { // Go won't create a copy of the original array if we only access the index
		fmt.Println(x)
	}

	for x := 0; x < len(books); x++ { // Better way to access if there is a big array
		fmt.Println(books[x])
	}

	// Create array from len of another array

	var published [len(books)]bool

	fmt.Printf("Published: %#v\n", published)

	// Ways to create arrays

	var books2 = [4]int64{1, 2, 3, 4}
	fmt.Printf("Books2: %#v\n", books2)

	// Use short declarations only if we know the values of our array

	// Faster way to create array
	books3 := [4]string{"Hola", "Como", "Estas", "Hoy?"}
	fmt.Printf("Books2: %#v\n", books3)

	books4 := [4]string{"Hola", "Como"} // Empty values are going to be initialized with 0 values , in the case of strings value is ""
	fmt.Printf("Books2: %#v\n", books4)

	newIntegers := [...]float32{1.2, 1.3, 1.4} // ... mean that is is going to create an array with length of 3 (values provided)
	fmt.Printf("Books2: %#v\n", newIntegers)

	// Compare arrays

	var farr [3]int64
	farr[0] = 21
	farr[1] = 22
	farr[2] = 23
	sarr := [...]int64{21, 22, 23}
	fmt.Println(farr == sarr) // true

	n := [4]int64{1, 2, 3, 4}
	n2 := [4]int64{2, 2, 3, 4}
	fmt.Println(n == n2) //Cannot compare for example n := [4]int64{1, 2, 3, 4} to n2 := [4]int32{2, 2, 3, 4} because of types (int32 , int64)

	// Assign array to another array

	testArr := [...]int32{3, 6, 9}
	copyArr := testArr //Type assigned in this case is [3]int32 type

	fmt.Println(copyArr)
	fmt.Println("Same?: ", testArr == copyArr) // true

	var copy2Arr [3]int32
	copy2Arr = testArr // Same type can be copied, cannot copy arrays from different types
	fmt.Println(copy2Arr)
	fmt.Println("Same?: ", testArr == copy2Arr) // true

	// Empty arrays
	var empty [3]uint8
	fmt.Println(empty)

	fmt.Println(strings.Repeat("~", 100))

	introArr := &testArr //Access to the memory itself
	introArr[0] = 10
	fmt.Println(introArr)
	fmt.Println(testArr)

	manyBooks := [...]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	upper, lower := manyBooks, manyBooks

	for i := range manyBooks {
		upper[i] = strings.ToUpper(upper[i])
		lower[i] = strings.ToLower(lower[i])
	}

	fmt.Printf("books: %q\n", manyBooks)
	fmt.Printf("upper: %q\n", upper)
	fmt.Printf("lower: %q\n", lower)

	// Multi-dimensional arrays, matrix

	students := [][]float64{
		{9, 6, 10},
		{7, 8, 6},
	}
	var sum float64
	var sum2 float64

	for x := 0; x < len(students[0]); x++ {
		sum += students[0][x]
	}

	for x := 0; x < len(students[1]); x++ {
		sum2 += students[1][x]
	}

	fmt.Println(sum, sum2)

	fmt.Println((sum + sum2) / float64(len(students)*len(students[0])))
	fmt.Println(students)
	fmt.Println(len(students[0]))

	const (
		ETH = iota
		WAN
	)

	// Keyed elements, makes everything more readable
	rates := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
		5:   11.1,
		2.6,
	}
	fmt.Println(rates)
	fmt.Println("1 BTC to ETH is :", rates[ETH])
	fmt.Println("1 BTC to WAN is :", rates[WAN])

	// Types in arrays

	//var green [5]int

	type integer int
	type bookcase [5]int
	type cabinet [5]int
	type anotherOne [5]integer
	orange := cabinet{1, 2, 3, 4, 5}
	yellow := bookcase{1, 2, 3, 4, 5}

	blue := [5]int{6, 9, 3, 2, 1}
	red := [5]int{6, 9, 3, 2, 1}
	fmt.Println(blue == red)
	fmt.Println(yellow == red) /* Difference here is that we can compare untyped with typed yellow (bookcase), red (no type)
	this only works if underlying type is the same, both are type int[5], we cannot do this with normal variables
	*/
	//fmt.Println(yellow == orange) cannot do this since types are the same, even when we have same underlying types
	// this same thing happens with normal variables

	green := anotherOne{1, 2, 3, 4, 5}
	fmt.Println(green)
	fmt.Println(green == [5]integer{1, 2, 3, 4, 5})

	//Solutions
	fmt.Println(yellow == bookcase(orange))
	fmt.Println(cabinet(yellow) == orange)
	fmt.Println([5]int(yellow) == orange)
	fmt.Println(yellow == [5]int(orange))

	fmt.Printf("Type %T\n", yellow)

	//Compare with variables
	type test1 int
	type test2 int

	purple := test1(5)
	white := test2(5)

	gray := 5
	var black int = 5
	fmt.Println(gray == black)
	fmt.Println(purple, white)
	//fmt.Println(white == black)
	/* Even if underlying types are equal, types are different so this cannot be done
	the same way we do with typed arrays, only untyped constants can do this
	*/
	fmt.Println(test2(black) == white)

	// can create array from constant
	const create int32 = 22
	var intros [create]int32

	fmt.Println(intros)
	fmt.Println(time.Now().Minute())

	zero := [5]string{
		"█████",
		"█   █",
		"█   █",
		"█   █",
		"█████",
	}

	for x:= range zero{
		fmt.Println(zero[x])
	}

}
