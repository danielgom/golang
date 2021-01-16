package main

import (
	"fmt"
	"reflect"
	"runtime"
)

var noLoSe = "talvez"

func main() {
	var hello = "goa"
	var dos int8 = 2.0
	fmt.Println(dos)
	fmt.Println(hello)
	fmt.Println("Hello!")
	bye()
	hey()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(reflect.TypeOf(hello))
	fmt.Println(reflect.TypeOf(dos))
	if 5 > 2 {
		fmt.Println("it is not null")
	}
	fmt.Println(goVersion())
	fmt.Println("\U0001f600")
	fmt.Printf(noLoSe)
	noLoSe = "asasd"
	fmt.Printf(noLoSe)

}
