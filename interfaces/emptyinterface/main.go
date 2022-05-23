package main

import "fmt"

func main() {

	// BE CAREFUL USING EMPTY INTERFACES , USE IT ONLY WHENEVER IT IS TRULY NECESSARY

	// Empty interfaces are implemented by definition, what it means is that anyInterface value implements this interface, however
	// as we can see, this interface does not have anyInterface method in it so we cannot use anything with anyInterface
	var anyInterface interface{}
	anyInterface = []int{1, 2, 3}
	anyInterface = map[string]bool{"ready": false, "set": false, "go": true}
	anyInterface = "this is a string"
	anyInterface = 22
	fmt.Println(anyInterface)
	// If we want to sum 22 with something , firstly we need to get the value from anyInterface and store it again
	anyInterface = anyInterface.(int) + 38 // We convert anyInterface to int then we sum 38 to that in and convert that value back to the anyInterface interface
	fmt.Println(anyInterface)

	nums := []int{1, 2, 3, 4, 5}
	anyInterface = nums
	fmt.Println(len(anyInterface.([]int))) // We need to convert anyInterface to the int slice before getting the length value of the slice

	var many []interface{} // Consist on many interface values
	for i := range nums {
		many = append(many, nums[i]) // Here we are appending the nums to the []interface called many, basically doing
		// the same thing as above and mapping every int value to the interface value just like we did with anyInterface
	}
	fmt.Println(many) // result: [1 2 3 4 5],
	fmt.Println(len(many))
	fmt.Println(many[0]) // Getting the number 1 of the many []interface

}
