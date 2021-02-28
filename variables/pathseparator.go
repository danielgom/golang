package main

import (
	"fmt"
	"github.com/dgomez/learngo/first/explain/exportlib"
	"github.com/dgomez/learngo/first/explain/library"
	"path"
	"strings"
)

func main() {

	var dirs, files string

	dirs, files = path.Split("asset/learning")
	fmt.Println("file: ", files)
	fmt.Println("dir: ", dirs)

	// faster less verbose
	dir, file := path.Split("asset/learning")
	fmt.Println("file: ", file)
	fmt.Println("dir: ", dir)
	fmt.Println(exportlib.GetGoVersion())
	library.PrintSomething()

	fmt.Println(strings.Fields("hola como estas"))

	res := Filter([]int{1, 2, 3, 4, 5}, func(a int) bool {
		return a%2 == 0
	})

	fmt.Println(res)

	next := func(a int) bool {
		return a%2 == 0
	}

	results := Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, next)
	fmt.Println(results[0:cap(results)])

}

func Filter(s []int, b func(a int) bool) []int {
	n := make([]int, 0, len(s)/2)
	for x := range s {
		if b(x) {
			n = append(n, x)
		}
	}
	return n
}
