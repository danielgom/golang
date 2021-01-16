package main

import "fmt"

type bottle struct {
	capacity float64
	height float64
}

func (b bottle) String() string {
	return fmt.Sprintf("The bottle has a capacity of %.2f, and a height of %.2f", b.capacity, b.height)
}
