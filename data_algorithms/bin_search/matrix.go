package main

import (
	"fmt"
)

func main() {
	var mat = [][]int{
		{1, 2, 3},
		{9, 10, 11},
		{21, 22, 23},
		{41, 56, 90},
	}

	b := exists(mat, 22)
	fmt.Println(b)
}

func exists(mat [][]int, number int) bool {
	firstP := 0
	rows := len(mat)
	cols := len(mat[0])
	secondP := (rows * cols) - 1

	for firstP <= secondP {
		midIdx := (firstP + secondP) / 2
		row := midIdx / cols
		col := midIdx % cols
		currentN := mat[row][col]
		if currentN > number {
			secondP = midIdx - 1
		} else if currentN < number {
			firstP = midIdx + 1
		} else {
			return true
		}
	}

	return false
}
