package main

import "fmt"

// Number interface example, this interface contains all types of number types
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// bubbleSort accepts Number interface it means that it can accept any type that is currently handled in Number interface
func bubbleSort[N Number](input []N) []N {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func newGenericFunc[age Number](myAge age) {
	fmt.Println(myAge)
	val := float32(myAge) + 1
	fmt.Println(val)
}

// Another example with types

type customFloat float64

// AddTwoFloats restricts all values that are not float64
func AddTwoFloats[T float64](a, b T) T {
	return a + b
}

// AddTwoFloatsOrSimilar accepts customFloat type since it is a "similar" type to float64
func AddTwoFloatsOrSimilar[T ~float64](a, b T) T {
	return a + b
}

func main() {

	var testAge = 28
	var testAge2 = 28.7

	list := []int32{4, 21, 5, 1, 7, 21, 2}
	listFloat := []float64{2.234, 63.23, 3.2, 5.6}

	sortedList := bubbleSort(list)
	sortedFloats := bubbleSort(listFloat)

	fmt.Println(sortedList)
	fmt.Println(sortedFloats)

	newGenericFunc(testAge)
	newGenericFunc(testAge2)

	custom := customFloat(22.22)
	custom2 := customFloat(54.123)

	// Error, customFloat is not a strict float64
	//AddTwoFloats(custom, custom2)

	fmt.Println(AddTwoFloats(12.4, 564.3))

	// No errors since custom and custom2 are similar types of float64
	fmt.Println(AddTwoFloatsOrSimilar(custom, custom2))

	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys map:", MapKeys(m))

	var m2 = map[int]int{1: 110, 2: 120}
	fmt.Println("Keys map2:", MapKeys(m2))
	linkedList := LinkedList[int]{}
	linkedList.Push(10)
	linkedList.Push(20)
	linkedList.Push(30)
	fmt.Println("LinkedList:", linkedList.GetAll())
	strLinkedL := LinkedList[string]{}
	strLinkedL.Push("hello")
	strLinkedL.Push("there")
	fmt.Println("String Linked List:", strLinkedL.GetAll())
}
