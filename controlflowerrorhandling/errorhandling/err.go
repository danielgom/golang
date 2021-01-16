package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {

	// Nil value is the null value in java
	// This is not for objects (Golang does not have objects)
	// This is meant for pointer types such as pointers, slices, maps , errors, etc ...

	fmt.Println("error handling")

	n, err := strconv.Atoi(os.Args[1])
	fmt.Println("no value")
	fmt.Println("Converted value: ", n)        // either returns expected value or 0 which means that an error happened
	fmt.Println("Returned error value: ", err) // <nil> means no error was returned (null) value
	var probe int64 = 22
	fmt.Println(math.Max(float64(probe), 34))

	// Handling error
	age, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Success: converted %q, to %d\n", os.Args[2], age)

	// Feet to meters
	feet := os.Args[3]
	feetInt, err := strconv.Atoi(feet)
	if err != nil {
		fmt.Printf("Error: %v is not a number\n", feet)
		return
	}
	meters := float64(feetInt) / 3.28084
	fmt.Printf("%d feet is %.2f meters.\n", feetInt, meters)

	// short if

	if a := os.Args; len(a) != 5 {
		fmt.Println("Give me a number please")
	} else if n, err := strconv.Atoi(a[4]); err != nil {
		// Variables seen : a , n , err
		fmt.Printf("Cannot convert %q. \n", a[4])
	} else {
		// All variables in scope (a ,b , err)
		fmt.Printf("%s * 2 is %d \n", a[4], n*2)
	}

	// switch statement there is no need for break statement after each one (like java switch expression)
	var city = ""
	switch city {
	case "Tokyo":
		fmt.Println("It is japan")
	case "Lyon", "France":
		fmt.Println("It is france")
	default:
		fmt.Println("None of them were chosen")
	}

	// Bool switch expression
	i := 10
	switch { // no need of a variable/value
	case i < 0:
		fmt.Println("negative")
	case i > 0:
		fmt.Println("positive")
	default:
		fmt.Println("It is not a number")
	}

	// Fallthrough statement
	i2 := 134
	switch {
	case i2 > 100:
		fmt.Println("big ")
		fallthrough // jumps the case and just prints positive
	case i2 > 0:
		fmt.Println("positive ")
		fallthrough // jumpst to print number
	default:
		fmt.Println("number")
	}

	// Short switch

	switch i:= 20; i {
	case 20:
		fmt.Println("Integer value is 20")
	case 30:
		fmt.Println("Integer vaule is 30")
	default:
		fmt.Println("None of the values found")
	}

	// Short switch boolean

	switch i:=30;{ // no need of declaring either true or false
	case i == 30:
		fmt.Println("Right now its value is 30")
		fallthrough
	case i > 10:
		fmt.Println("Value is higher than 10")
	default:
		fmt.Println("No value found in this switch")
	}

	fmt.Println(time.Now().Date())
}
