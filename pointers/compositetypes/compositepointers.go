package main

import (
	"fmt"
	"strings"
)

type house struct {
	name string
	rooms int
}

func main() {
	fmt.Println("Working with arrays")
	arrays()
	fmt.Println()
	fmt.Println()
	fmt.Println("Working with slices")
	slices()
	fmt.Println()
	fmt.Println()
	fmt.Println("Working with maps")
	maps()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Working with structs")
	structs()

}

func arrays() {
	nums := [...]int{1, 2, 3}
	fmt.Printf("Address of the array which is passed: %p\n", &nums)
	fmt.Println("Before calling the array normal function:", nums)
	incr(nums)
	fmt.Println("After calling the array normal function:", nums)

	fmt.Println()
	fmt.Println()

	fmt.Printf("Address of the array which is passed : %p\n", &nums)
	fmt.Println("Before calling the pointer function:", nums)
	incrByPointer(&nums) // changes the value of the array
	fmt.Println("After calling the pointer normal function:", nums)
}

func slices() {
	dirs := []string{"up", "down", "left", "right"}
	fmt.Printf("Address of the slice which is passed: %p\n", &dirs)
	fmt.Println("Before calling the slice function:", dirs)
	up(dirs)
	fmt.Println("After calling the slice function:", dirs)
	fmt.Println("Adding to the original slice in the function slice")
	modSlice(&dirs)
	fmt.Println("After calling the slice function (pointer function add):", dirs)
	fmt.Println("DO NOT WORK WITH POINTERS AND SLICES")

}

func maps() {
	confused := map[string]int{"one": 1, "two": 2}
	fmt.Println("Before calling the map function:", confused)
	fix(confused)
	fmt.Println("after calling the map function:", confused)
	fmt.Println("DO NOT WORK WITH POINTERS AND MAPS")

}

	// NEED TO BE CAREFUL WHEN USING STRUCTS AND POINTERS SINCE WE MAY SEE SOME CHANGES IN THE SLICES AND MAPS WE HAVE
	// WITHIN THE STRUCT

func structs() {
	myHouse := house{name: "My House", rooms: 5}
	fmt.Println("Before calling function on struct:", myHouse)
	addRoom(myHouse)
	fmt.Println("After calling function on struct:", myHouse)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Before calling function on struct (pointer function):", myHouse)
	addRoomPointer(&myHouse)
	fmt.Println("After calling function on struct (pointer function):", myHouse)

}

func incr(nums [3]int) { // This function creates another array and every change made here is not noticeable in the caller
	fmt.Printf("Address of the array which is created in the function called: %p\n", &nums)
	for x := range nums {
		nums[x] ++
	}
}

func incrByPointer(nums *[3]int) {
	fmt.Printf("Address of the array which is created in the function called (pointer function): %p\n", &nums)
	for i := range nums {
		nums[i]++
	}
}

func up(slice []string) { // Even if the address of both slices (original and this one in the function) are different
	// golang is pointing to the slice header so any change here is going to affect the original slice
	// EXCEPT APPENDING, APPENDING REQUIRES TO RETURN THE APPENDED SLICE
	fmt.Printf("Address of the slice within the function: %p\n", &slice)
	for i := range slice {
		slice[i] = strings.ToUpper(slice[i])
	}
	slice = append(slice, "newline") // Does not work since this is a copy slice with a different backing array
	// this can only work in the next function which is using pointers with slices
}

func modSlice(il *[]string) {
	*il = append(*il, "newline")
}

func fix(m map[string]int) {
	m["three"] = 3
}

func addRoom(h house) {
	h.rooms++
}

func addRoomPointer(h *house) {
	h.rooms++
}
