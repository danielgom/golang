package main

import (
	"fmt"
	"github.com/danielgom/go_math/calc"
	"github.com/danielgom/go_math/geometry"
)

func main() {

	/*
		We are going to be using the module that is placed on /Users/danielg/go_modules/go_math
		Url used https://github.com/danielgom/go_math
		for this we did  "go get github.com/danielgom/go_math/calc"
	*/

	add := calc.Add(20, 30)
	fmt.Println(add)
	volume := geometry.CubeVolume(50)
	fmt.Println(volume)
}
