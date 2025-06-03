package main

import (
	"fmt"
	"github.com/samber/lo"
	"slices"
)

func main() {

	fmt.Println("hello there")

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	newSl := lo.Filter(mySlice, func(myVal int, _ int) bool {
		return myVal%2 == 0
	})

	fmt.Println(len(newSl))
	fmt.Println(cap(newSl))
	fmt.Println(newSl)

	mySecondSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newSl2 := slices.DeleteFunc(mySecondSlice, func(myVal int) bool {
		return myVal%2 != 0
	})

	fmt.Println(len(newSl2))
	fmt.Println(cap(newSl2))
	fmt.Println(newSl2)

}
