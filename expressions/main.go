package main

import (
	"fmt"
	"github.com/dgomez/learngo/first/explain/library"
	"reflect"
	"runtime"
)

/*
Main function executable commands one only main function for
Executable programs
*/
func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.Version())
	doNow()
	var nope = "ho"
	fmt.Println(reflect.TypeOf(nope))
	fmt.Println(nope[1])
	library.PrintSomething()
	fmt.Println(library.Number + 39)
	fmt.Println("Go lang version is: ", library.ReturnGoVersion())
	// this is a comment
}

//doNow prints something amazing
func doNow() {
	fmt.Println("this just prints something")
	fmt.Println("inanc")
	fmt.Println("lina")
	fmt.Println("ebru")
}
