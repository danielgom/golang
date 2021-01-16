package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		start rune
		stop  rune
	)

	start, stop = 'A', 'Z'
	//Or
	// start, stop := 'A','B' Without the var creation
	fmt.Println(start, stop)

	for x := start; x <= stop; x++ {
		fmt.Printf("%-10c %[1]d\n", x)
	}

	fmt.Println("Separation")
	fmt.Println()
	fmt.Println()
	fmt.Println()

	var begining, end int

	if args := os.Args[1:]; len(args) == 2 {
		begining, _ = strconv.Atoi(args[0])
		end, _ = strconv.Atoi(args[1])
	}

	if start == 0 || stop == 0 {
		start, stop = 'A', 'Z'
	}

	for x := begining; x <= end; x++ {
		fmt.Printf("%-10c %-10[1]d % -12x\n", x, string(x))
	}

	// cant
	//var news byte = '😂' Causes overflow since 😂 is represented by 128514 and cannot be stored in a byte
	//fmt.Println(news)

	// Can
	var news = '😂'	// Rune can save up to a 4 byte representation character , can save 128514
	//Remember that rune = int32
	fmt.Println(news) // Prints 128514
	fmt.Printf("%c", news)		// Prints the emoji
}
