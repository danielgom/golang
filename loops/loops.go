package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const max = 5

func main() {

	var tt = 2
	tt++
	next := tt
	fmt.Println(next)
	fmt.Println("Loops in golang")
	var sum int

	for i := 0; i <= 5; i++ {
		sum += i
	}
	fmt.Printf("Sum: %d\n", sum)

	// infinite for

	i := 0
	for { // always true
		if i >= 5 {
			break // need to break to exit the infinite for loop
		}
		fmt.Println("Still on the loop")
		i++
	}

	// continue same as the other thing
	// just jumps to the next for

	x := 0
	for { // always true
		if x >= 5 {
			break
		}
		if x == 2 || x == 3 {
			x++
			continue
		}
		fmt.Println("Still on the loop, number: ", x)
		x++
	}

	// print multiplication table
	fmt.Printf("%5s", "X")
	for h := 0; h <= max; h++ {
		fmt.Printf("%5d", h)
	}

	fmt.Println()

	for x := 0; x <= max; x++ {
		fmt.Printf("%5d", x)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", x*j)
		}
		fmt.Println()
	}

	fmt.Println()

	tableSize, err := strconv.Atoi(os.Args[1])
	if err != nil || tableSize <= 0 {
		fmt.Println("Wrong size")
	}

	fmt.Printf("%5s", "X")
	for h := 0; h <= tableSize; h++ {
		fmt.Printf("%5d", h)
	}

	fmt.Println()

	for x := 0; x <= tableSize; x++ {
		fmt.Printf("%5d", x)

		for j := 0; j <= tableSize; j++ {
			fmt.Printf("%5d", x*j)
		}
		fmt.Println()
	}

	// Slice

	words := strings.Fields("lazy cats jumps too many times that i'm getting exhausted")
	fmt.Println(words[0])

	for i := 0; i < len(words); i++ {
		fmt.Println("Index: " + strconv.Itoa(i) + " word: " + words[i])
	}

	for index, value := range words {
		fmt.Println("Index: ", strconv.Itoa(index), " word: "+value)
	}

	// if only index needed , same happens with values
	for index := range words {
		fmt.Println("Index: ", strconv.Itoa(index))
	}

	for i := 3; i > 0; {
		i--
		fmt.Println(i)
	}

	for {
		var c string

		switch rand.Intn(4) {
		case 0:
			c = "\\"
		case 1:
			c = "/"
		case 2:
			c = "|"
		case 3:
			c = "-"
		}
		fmt.Printf("\r%s Please Wait. Processing....", c)
		time.Sleep(time.Millisecond * 250)
	}
}
