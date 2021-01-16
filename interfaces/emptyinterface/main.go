package main

import "fmt"

func main() {

	// BE CAREFUL USING EMPTY INTERFACES , USE IT ONLY WHENEVER IT IS TRULY NECESSARY

	// Empty interfaces are implemented by definition, what it means is that any value implements this interface, however
	// as we can see, this interface does not have any method in it so we cannot use anything with any
	var any interface{}
	any = []int{1, 2, 3}
	any = map[string]bool{"ready": false, "set": false, "go": true}
	any = "this is a string"
	any = 22
	fmt.Println(any)
	// If we want to sum 22 with something , firstly we need to get the value from any and store it again
	any = any.(int) + 38 // We convert any to int then we sum 38 to that in and convert that value back to the any interface
	fmt.Println(any)

	nums := []int{1, 2, 3, 4, 5}
	any = nums
	fmt.Println(len(any.([]int))) // We need to convert any to the int slice before getting the length value of the slice

	var many []interface{} // Consist on many interface values
	for i := range nums {
		many = append(many, nums[i]) // Here we are appending the nums to the []interface called many, basically doing
		// the same thing as above and mapping every int value to the interface value just like we did with any
	}
	fmt.Println(many) // result: [1 2 3 4 5],
	fmt.Println(len(many))
	fmt.Println(many[0]) // Getting the number 1 of the many []interface


}
