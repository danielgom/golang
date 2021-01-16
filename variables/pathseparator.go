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

}
