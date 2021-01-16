package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jump again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries: //Labeled break, means that it will break from the next for it encounters(Continue is able to do this as well)
	for _, x := range query {
		for in, j := range words {
			if x == j {
				fmt.Printf("#%-2d: %q\n", in+1, j)
				break queries
			}
		}
	}

	fmt.Println(strings.ToLower("HELLO"))
}
