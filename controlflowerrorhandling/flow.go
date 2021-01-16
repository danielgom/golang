package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Bolean examples
	var nope bool

	nope = 20 >= 22.3
	fmt.Println(nope)

	speed := 100
	fast := speed >= 80
	slow := speed < 20

	fmt.Println(fast, slow)

	// Need to compare same types or typeless values

	speedB := 150.5
	faster := speedB > 80 // 80 is typeless so it is converted to float64
	if faster {
		fmt.Println("it is fast")
	}

	//Main if example
	parseInt, _ := strconv.ParseInt(os.Args[1], 10, 64)
	if parseInt > 20 {
		fmt.Println("number typed is greater than 20")
	} else {
		fmt.Println("number typed is not grater than 20")
	}

	// && And operator
	// || Or operator
	// ! not Operator ... !false == true ... !true == false

	// else if
	score := parseInt
	if score > 3 {
		fmt.Println("Score is greater than 3")
	} else if score == 21 {
		fmt.Println("Score is 20")
	} else {
		fmt.Println("Score is lesser than 20")
	}

	message := `
This is a good message
the main problem here is that there is a white space`

	if len(os.Args) < 3 {
		fmt.Println(strings.TrimSpace(message))
		fmt.Println()
		fmt.Println("this is space")
		fmt.Println(strings.TrimSpace(message))
	}
	fmt.Println(strings.Index("a", "aeiou"))           // -1 mean that none is found (need to find the whole substring)
	fmt.Println(strings.Index("a a b aeiou", "aeiou")) // 6 first index found in the whole substring
	fmt.Println(strings.Contains("a", "aeiou"))        // false, it means that the whole substring is not found
	fmt.Println(strings.IndexAny("a", "aeiou"))        // 0 means it has found it in the index 0 (finds if any char is on the main string )
	fmt.Println(strings.ContainsAny("a", "aeiou"))     // true , means that it has found any char in the main string

}
