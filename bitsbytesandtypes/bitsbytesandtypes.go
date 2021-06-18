package main

import (
	"fmt"
	"github.com/dgomez/learngo/first/explain/numbersandstrings/weights"
	"math"
	"strings"
	"time"
)

type gram float64
type ounce float64

func main() {
	fmt.Printf("%04b", 3)

	//var testSigned int8 = 128 // cant do this signed
	var test uint8 = 128 // can, since it is unsigned
	fmt.Println(test)
	var nope int = 23
	fmt.Println(math.MaxInt8 + nope)

	f := math.MaxFloat32
	fmt.Println(math.Pow10(1))
	fmt.Println(math.Pow(1, 1))
	fmt.Println(f * 1.2)

	sentence := "anita lava la tina"
	fmt.Println(strings.ReplaceAll("", sentence, "asdasd"))

	pre := 12
	fmt.Println(float64(pre))
	var m int64 = 2
	hours, _ := time.ParseDuration("4h30m")
	hours *= time.Duration(m)
	fmt.Println(hours)

	// using types underlying types

	var g gram = 1000
	var o ounce
	//o = g * 0.35742 //cant since g is gram and o is ounce
	o = ounce(g) * 0.03574 // need to convert
	fmt.Printf("%g grams in %.2f ounces\n", g, o)

	type (
		Gram     int
		Kilogram Gram
		Ton      Kilogram
	)

	var salt Gram = 100
	var apple Kilogram = 5
	var truck Ton = 10

	g2 := Gram(2)
	fmt.Println(g2)

	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apple, truck)

	salt = Gram(apple)
	truck = Ton(salt)
	fmt.Println(salt, truck)
	interestingVariable := Ton(Gram(int(apple)))
	fmt.Println(interestingVariable)
	myGram := weights.Kilogram(10)
	fmt.Println(myGram)
	fmt.Println(time.Duration(10))

	type MyAnotherGram weights.Gram
	var grammsss = MyAnotherGram(weights.Gram(10))
	sumInt := Ton(weights.Kilogram(10)) + Ton(grammsss)
	fmt.Println(sumInt)

	type Millennium time.Duration
	millennium := Millennium(10)
	fmt.Println(int64(millennium))

	type Distance int

	const (
		village Distance = 50
		city             = 100
	)

	fmt.Print(village + Distance(city))
	fmt.Println(time.Duration(2) + 20)
	word := "hello"
	fmt.Println(strings.ReplaceAll(word, "llo", "nope"))
	fmt.Println("hello world")
	fmt.Println("check for terminal")

	//var talvez int64
	type probar int64
	type probando probar
	var exit probar = 22
	var exit2 probando = 23
	//fmt.Println(exit + exit2) // cannot do this, since these are different
	fmt.Println(int64(exit) + int64(exit2))
	//fmt.Println(exit == talvez)

	testSlice := []string{"hola", "como", "estas", "?"}
	fmt.Println(testSlice)

	var testInts []int
	fmt.Println(testInts)
	testInts = append(testInts, []int{1, 2, 3, 4}...)
	fmt.Println(testInts)

	testSplit := testInts[0:3]
	fmt.Println(testSplit)
	fmt.Println(cap(testSplit))
	fmt.Println(testSplit[0:cap(testSplit)])


	testSplit[0] = 22
	fmt.Println(testInts)


	testCopy := make([]int, len(testInts))
	copy(testCopy, testInts)
	fmt.Println(testCopy)
	testCopy[1] = 33
	fmt.Println(testInts)
	fmt.Println(testCopy)


}
