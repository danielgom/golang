package main

import (
	"flag"
	"fmt"
	"reflect"
)

func main() {
	speed, force := 100, 2.5

	args := flag.Args()
	fmt.Println("here are the args", args)
	speed = int(float32(speed) * float32(force))
	fmt.Println(speed)
	velocity := 22
	fmt.Println(velocity)
	fmt.Println(reflect.TypeOf(velocity))
}
