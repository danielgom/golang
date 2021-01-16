package main

import (
	"fmt"
	"github.com/dgomez/learngo/first/explain/library"
	"os"
)

func main() {

	var prime, second = "hola", 123

	second, third := 43, 1.24243

	fmt.Println(prime, second, third)

	fmt.Println(float64(second))

	library.PrintSomething()
	arguments := os.Args[2]
	fmt.Println(arguments[1], arguments[2], "This is the first element")
	fmt.Println(os.Args)
	fmt.Println(os.Args[0], "Printing the path of the program")
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])
	fmt.Println(len(os.Args))
}
