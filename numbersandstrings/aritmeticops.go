package main

import (
	"fmt"
	"math"
	"os"
	"os/user"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	var average float64
	myAge, yourAge := 22, 33
	newAverage := myAge + yourAge
	average = float64(myAge+yourAge) / 2
	fmt.Println(average)
	fmt.Println(reflect.TypeOf(average), reflect.TypeOf(newAverage))

	// bad precision

	ratio := 1.0 / 10
	fmt.Printf("%.60f\n", ratio)
	fmt.Printf("%.60f\n", ratio*10.0)

	fmt.Println("sum : ", 3+2)

	// division
	var result float64 = 3 / 2
	fmt.Println(result) // converts into ints
	result = float64(int(3) / int(2))
	fmt.Println(result)

	// solution

	result = 3.0 / 2
	fmt.Println(result)
	result = float64(3) / 2
	fmt.Println(result)

	myAge++
	fmt.Println(myAge)

	// small exercise
	// celcius to farenheit

	var training float64 = 2 * 2
	fmt.Println(training)
	celsius, _ := strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("Celsius %f equals to %g farenheit\n", celsius, celsius*1.8+32)
	fmt.Println(math.Hypot(3, 4))
	fmt.Println(math.Sqrt(32))
	fmt.Println(math.Pow(32, 2))

	checkString := ""
	fmt.Println(len(checkString))

	// String vs raw String ....

	var testString string
	testString = "hello\n" + // need jump line
		"world"
	fmt.Println(testString)

	// Raw string does not need jumps
	rawString := ` 
<html>
	<body>"Hello"</body>
</html>`
	fmt.Println(rawString)

	// Concat works the same for both
	fmt.Println("Hello "+"how are you? ", `today I'm not feeling
very good`, "and you?\n", `I'm fine`)

	// Convert from any other type to string
	eq, sum := "1 + 2 = ", 3
	fmt.Println(eq + strconv.FormatInt(int64(sum), 10))
	fmt.Println(eq + strconv.Itoa(sum)) // same as above

	// Length of a String
	exampleString := "∫∂letterœ"
	fmt.Println(len(exampleString)) // prints 14 but it has only 9 letters

	fmt.Println(utf8.RuneCountInString(exampleString)) // prints correctly 9 letters (which are 9 runes)
	fmt.Println(utf8.DecodeRuneInString(exampleString))
	fmt.Println(utf8.FullRuneInString(exampleString))
	current, _ := user.Current()
	fmt.Printf("%v\n", current)

	// Strings package
	fmt.Println(strings.Repeat("!", 5))
	fmt.Println(os.Args[1] + strings.Repeat("!", 5))
	fmt.Println("2" + "hola")
	repeat := strings.Repeat("!", 3)
	fmt.Println(repeat + os.Args[1] + repeat)
	fmt.Println(strings.ToUpper("daniel"))
	fmt.Println(utf8.RuneCountInString("hétérogénéité"))

	name := "inanç           "

	name = strings.TrimRight(name, " ")
	fmt.Println(name)
	l := utf8.RuneCountInString(name)

	fmt.Println(l)

	msg := `
	
	The weather looks good.
I should go and play.
	`

	fmt.Println(msg)
	fmt.Println(strings.TrimSpace(msg))
}
