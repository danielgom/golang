package main

import "fmt"

// If we want to use the []*game in a method we should declare another type variable and use it as we are using it in the
// following code
type list []*game

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("There are no elements on the list")
		return
	}
	for i := range l {
		l[i].print()
	}
}
/*
func (g []*game) name() { (error) we cannot do something like this since our game type is declared in another file

}

 */