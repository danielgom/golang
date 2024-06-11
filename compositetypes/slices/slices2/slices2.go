package main

import (
	"fmt"
	"sort"
)

type collection [5]int
type slcollection []int

func main() {

	// Continuation of slices/slices

	/*
		In slices, there is a backing array so the following code is going to test how it completely works under
		the hood
	*/

	grades := []float64{20, 30, 10, 40, 50, 60}

	/*
		it does not create a new slice, this line of code just refers to the numbers
		{20,30,10}, under the hood a slice has an array of elements , each slice created from this slice
		points to the same array
	*/
	sl1 := grades[:3]
	sl2 := grades[1:4]
	fmt.Println(sl1)

	/*
		Capacity Example:
		Sl1 has capacity of 6 because it starts from the 0 index to the 3 element, however capacity stands for
		how many elements of the backing array can be seen , in this case we can see the 6 elements

		Sl2 has a capacity of 5 because it can se values from 30 to 60 in the backing array, it cannot look back to
		the number 20, so capacity is 5

	*/

	fmt.Println(cap(sl1)) // Here capacity is 6
	fmt.Println(cap(sl2)) // Here capacity is 5

	/*
		If we change something in the slice sl1, it is going to affect grades slice since both of them are
		pointing to the same array which is under the hood
	*/

	/*
		Sorting the slice is going to sort the array under the hood {20,30,10} (in place sorting)
	*/

	sort.Float64s(sl1)
	fmt.Println(sl1)

	// Printing the grades slice is going to show the first three elements sorted since we sorted it in the sl1 slice
	fmt.Println(grades)

	/*
		Appending does not work the same way, it creates a new slice
	*/

	newGrades := grades
	fmt.Println(newGrades)

	moreGrades := []float64{70, 80, 90, 100}
	newGrades = append(newGrades, moreGrades...) // creating a new slice from append
	fmt.Println(newGrades)                       // Original 6
	fmt.Println(grades)                          // new Slice with more numbers

	// Difference when passing an array and a slice to a function (remember slice is a reference to a backing array)
	//fmt.Println(fibon(5))
	//fmt.Println(fibor(5))

	/*
		When passing array to a function, this function creates a copy array from the array that is passed
		so if we have 1 million values in an array and pass into that function, then that function copies
		all values within that array and does its thing (costly)

		Slices, a slice passed to a function does not copy a slice, it copies the slice header which contains
		the information of the backing array, so if we change something then the original slice is going to be
		changed as well (not costly operation)
	*/

	check := collection{1, 2, 3, 4, 5}

	modify(check)

	fmt.Println("Original value (after using function): ", check)

	checkSlice := slcollection{1, 2, 3, 4, 5}

	slmodifiy(checkSlice)

	fmt.Println("Original value (after using function, slice): ", checkSlice)

	fastTest := []int{5, 10, 15, 20, 25} // Cap 5, Len 5, pt ...
	someF := fastTest[1:3]               // Cap 4 (starting from 10 to end of backing array which is 25), len 2, pt ...
	fmt.Println(someF)
	fmt.Println(len(someF))
	fmt.Println(cap(someF))
	someF = append(someF, 10, 100, 200) // Crazy: adding this surpasses its own capacity, which means surpassing backing array's capacity, therefore another
	// backing array is created, if we do not surpass the capacity, then appending just replaces values
	// If we surpass the backing array, then it means that we are going to be creating a ned backing array
	// with double the capacity of itself, in this case
	someF[2] = 55
	fmt.Println(someF)
	fmt.Println(cap(someF))         // Capacity here is double the capacity someF had (4), then it is going to be 8
	fmt.Println(someF[:cap(someF)]) // we are here expanding the array to its own capacity, remember last line it doubled it so new values are going to be 0
	fmt.Println(len(someF))
	fmt.Println()
	fmt.Println(cap(fastTest))
	fmt.Println(fastTest)

	fastTest2 := []int{5, 10, 15}
	fmt.Println(cap(fastTest2))
	fastTest2 = append(fastTest2, 100)      // BETTER EXAMPLE
	fmt.Println(cap(fastTest2))             // Double its own capacity, now its 6, but length is 4
	fmt.Println(len(fastTest2))             // Len is 4
	fmt.Println(fastTest2[:cap(fastTest2)]) // Resized to its own capacity, this did not happen in place so we just print the whole thing {5,10,15,100,0,0}
	fastTest2 = fastTest2[:cap(fastTest2)]  // Resized and saved into its own array
	fmt.Println(len(fastTest2))             //
	fmt.Println(cap(fastTest2))
	fmt.Println(fastTest2[len(fastTest2):])
	fmt.Println(fastTest2)

	words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
	words = append(words[:1], "is", "everywhere")
	words = append(words, words[len(words)+1:cap(words)]...)
	fmt.Println(words)

	words2 := []string{"osi", 1022: "", 12: "hola", "nolose", "talvez"} // Works the say way as arrays, checks values and puts them on the right order
	words2 = append(words2, "boom!")
	fmt.Printf("%#v\n", words2)

	// doubles (cap is 1023, hence < 1024)
	fmt.Println("doubles:")

	wordss := []string{1022: ""}
	fmt.Println("l:", len(wordss), "c:", cap(wordss))
	wordss = append(wordss, "boom!")
	fmt.Println("l:", len(wordss), "c:", cap(wordss))

	// ------------------------------------------------

	// grows 1.25 (cap is 1024, hence >= 1024)
	fmt.Println("\ngrows %25:")

	wordss = []string{1023: ""}
	fmt.Println("l:", len(wordss), "c:", cap(wordss))
	wordss = append(wordss, "boom!")
	fmt.Println("l:", len(wordss), "c:", cap(wordss))

	/*
		Based on test done before, appending affects the main slice only if the slice that we are appending to, can still
		see the main slice, there are some comments above
	*/

	// Here we are going to see how for loop affect slices

	testFor := []string{"Hello", "this", "is", " a ", "for", "test"}
	testFor2 := testFor

	for _, v := range testFor { // When we loop over a slice with a for loop and we use the value, it clones
		// the backing array so the original slice takes no effetc
		fmt.Println(v)
		v = "empty"
		fmt.Println(v)
	}
	fmt.Println(testFor)  // same as before  [Hello this is  a  for test]
	fmt.Println(testFor2) // [Hello this is  a  for test] (copied from testFor)

	// Appending two slices, one non-nil and nil

	myFirstSlice := []int{1, 2, 3}
	var mySecondSlice []int

	fmt.Println(append(myFirstSlice, mySecondSlice...))

	var myFirstSliceNil []int
	var mySecondSliceNil []int

	fmt.Println(myFirstSliceNil == nil)
	fmt.Println(mySecondSliceNil == nil)
	fmt.Println(append(myFirstSliceNil, mySecondSliceNil...) == nil)
	fmt.Println(append(myFirstSliceNil, mySecondSliceNil...))
}

/*
func fibon(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 3 {
		return 1
	}
	actual := 1
	last := 1
	for x := 4; x <= n; x++ {
		temp := actual
		actual = actual + last //2
		last = temp	 // 1 +1 // 2
	}
	return actual
}

func fibor(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 3 {
		return 1
	}
	return fibor(n - 2) + fibor(n - 1)

*/

func modify(data collection) {
	data[0] = 10
	data[1] = 20
	data[2] = 30
	fmt.Println("Within the function: ", data)
}

func slmodifiy(data slcollection) {
	data[0] = 10
	data[1] = 20
	data[2] = 30
	fmt.Println("Within the function (slice): ", data)
}
