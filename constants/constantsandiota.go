package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	// constant goes in compile time whereas variables are runtime
	n := 10
	const word string = "hello"

	fmt.Println(word)

	//const value int = n // cant , n is not a constant,
	//const value int = int(math.Pow10(10)) // cannot assing functions to constants

	fmt.Printf("Must use the variable n: %d\n", n)
	// Break the rule , len
	const value = len("this is constant")
	fmt.Println(string("hello"[1] - 32))
	fmt.Println(strings.Index("hello", "h"))

	char := 'c' // rune
	char2 := "a"
	fmt.Println(char, char2)

	fmt.Printf("value type: %T\n", value)

	const min, max int = 5, 10
	// or

	const (
		min2 int = 10
		max2 int = 20
		noExpression
		noExpression2
	)

	fmt.Println(noExpression, noExpression2)

	var uno = 11 + 11
	var dos = 3.14
	fmt.Println(dos * float64(uno)) // need to convert with variables

	const first = 11 + 11          // typeless
	const secondens = 3.14         // typeless
	fmt.Println(first * secondens) // can be done without conversion since these are typeless

	const prime = 42 // can assign to integer since it is typeless
	var i float64
	i = 10 + prime // HERE GOLANG CONVERTS TO THE FLOAT64 TYPE BEHIND THE SCENES (IT HAPPENS WITH EVERY TYPELESS CONSTANT)
	fmt.Println(i)

	const mins int32 = 1
	maxim := 5 + mins
	fmt.Printf("Type: %T\n", maxim)
	fmt.Println(maxim)

	var probe = 24
	//fmt.Println(probe + 3.123) // cannot do this

	x := 32 + 4.5
	fmt.Println(x + float64(probe)) // need to convert

	// Constants with types
	type MyDuration int64
	const (
		NanoSecond  MyDuration = 1
		MicroSecond            = 1000 * NanoSecond // Converts to Nanosecond
		MilliSecond            = 1000 * MicroSecond
		Second                 = 1000 * MilliSecond
		Minute                 = 60 * Second
		Hour                   = 60 * Minute
	) // All types are "MyDuration"

	var duracion MyDuration = 22

	fmt.Println(duracion * Minute)

	t := time.Second * 10 // Convert 10 to time.Second since 10 is typeless
	fmt.Println(t)
	isd := 10
	//fmt.Println(time.Second * isd) // Cannot convert isd to time.Second since isd is int
	fmt.Println(time.Second * time.Duration(isd)) //Solution, convert to Duration

	const c = 10                 // Constant typeless
	const c2 int = 10            // Non-Typeless constant
	fmt.Println(time.Second * c) // Can do it since c is constant typeless
	//fmt.Println(time.Second * c2) // cant since c2 even if it is constant it is not typeless

	// small program

	const (
		feetInMeters = 0.3048 // Default value is float64
		feetInYards  = feetInMeters / 0.9144
	) // Only feetInYards is typeless

	var socabon = 22
	//last := socabon * feetInYards // int because feetinYards is typeless
	fmt.Println(math.Round(float64(socabon)))
	fmt.Println(time.January)

	// Iota number generator examples

	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	const (
		EST = -(5 + iota) // -5
		_                 // -6
		MST               // -7
		PST               // -8
	)
	fmt.Println(EST, MST, PST)

	// Naming conventions in Golang
	fmt.Println(time.Duration(1))

	nonRep := []int{1, 2, 3, 4, 5, 5, 5}
	fmt.Println(nonRep)
}
