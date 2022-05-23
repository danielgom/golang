package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	// ******************APPENDING APPENDS TO THE LAST LENGTH VALUE OF AN ARRAY/SLICE*****************************

	/*
		Having done everything so far, how do we prevent to erase the backing array with append
	*/

	sliceable := []int{1, 2, 3, 4, 5}
	half := sliceable[0:3]
	fmt.Println("Sliceable: ", sliceable)              //
	fmt.Println("Sliceable length ", len(sliceable))   // length 5
	fmt.Println("Sliceable capacity ", cap(sliceable)) // Backing array is 5

	fmt.Println("Half: ", half)
	fmt.Println("Half length: ", len(half)) // len is 3 since we can only get 3 numbers [0:3]
	fmt.Println("Half capacity", cap(half)) // This can see the whole backing array so its 5

	half = append(half, 100, 200) // With this, since we have capacity of 5, we erase 4 and 5 from the backing
	// array with 100 and 200
	fmt.Println("Half after appending: ", half)         // [1 2 3 100 200]
	fmt.Println("Sliceable after appending", sliceable) // [1 2 3 100 200]

	// If we were to prevent this we know that we must add another value so it exeeds capacity and creates another slice
	// without affecting the backing array of the sliceable variable

	// Better fix, slice expressions

	sliceable2 := []int{10, 20, 30, 40, 50}
	half2 := sliceable2[0:3:3] // Notice the last number, it means that capacity is going to be capped at
	// 3, no longer we would be able to see the 5 numbers of the backing array
	// this acts like a hide for 40,50 respectively

	fmt.Println("Sliceable2: ", sliceable2)              //
	fmt.Println("Sliceable2 length ", len(sliceable2))   // length 5
	fmt.Println("Sliceable2 capacity ", cap(sliceable2)) // Backing array is 5

	fmt.Println("Half2: ", half2)
	fmt.Println("Half2 length: ", len(half2)) // len is 3 since we can only get 3 numbers [0:3]
	fmt.Println("Half2 capacity", cap(half2)) // This is the great change

	fmt.Println("Half2 after appending: ", half2)         // [10 20 30]
	fmt.Println("Sliceable2 after appending", sliceable2) // [10 20 30 40 50] since we saw that capacity is 3 and our length was 3
	// if we append what is going to happen is that it is going to create a new backing array
	// this wont affect sliceable 2 backing array

	testing := []string{"Can", "You", "See", "Me", "?"}
	breaking := testing[0:3]                                                             // Capacity is still 5, no slice expression
	fmt.Println(testing)                                                                 // [Can You See Me ?]
	fmt.Println(breaking)                                                                // [Can You See]
	fmt.Println("Expanding to the backing array's capacity: ", breaking[:cap(breaking)]) // [Can You See Me ?]
	// This is because like we said, capacity is still 5 so we can expand til 5 and see all the values in the backing array

	// make function in action

	// MAKE FUNCTION CAN CREATE SLICES, MAPS AND CHANNELS

	lowerList := []string{"this", "is", "me", "again"}

	upperList := make([]string, 0, cap(lowerList))
	/*
		Basically , make creates an slice with length 0 and capacity of the capacity in the loweList slice which is 4
		all the values in the slice are 0 values, in this case since it is a string values are "", if we append, since we start
		from length 0 then it is going to append to the 0 element
	*/

	fmt.Println(len(upperList)) // len 0
	fmt.Println(cap(upperList)) // capacity 4

	for x := range lowerList {
		upperList = append(upperList, strings.ToUpper(lowerList[x]))
	}

	fmt.Println(upperList)

	testn2 := []string{"maybe", "this", "is", "not", "a", "good", "idea"}
	change := make([]string, 2, cap(testn2)) // length 2 capacity 7, then we can append upt to 5
	fmt.Println(cap(change))
	fmt.Println(cap(append(change, "", "", "", "", "", ""))) // Here we appended more than 5 so it creates a new backing array which
	// capacity is 14

	change2 := make([]int, 3)     // Main problem here is that we start with an array of length 3 and capacity 3
	change2 = append(change2, 10) // since length is at 3 and capacity is 3, then its going to append to the number after length which is 4
	// therefore it is going to create a new backing array with 6 of capacity
	// instead of appending there is a solution
	//change2[0] = 10			// Indicating index 0 would mean that changes position 0 same as having declared length 0 an another capacity
	// but this is not the solution we want
	fmt.Println(cap(change2))

	// Copy arrays
	// There are many ways to copy arrays, this is one of them
	un := []float64{1.2, 3.5, 124.2, 453.112}

	sec := make([]float64, len(un)) /// Make also creates a new slice (already known)

	copy(sec, un)    // First we copy , this happens in place not like append
	fmt.Println(sec) // Same output as

	// Both of them have different backing arrays, lets check what happens if we try to copy to a slice that has a lower capacity

	bigCap := []int{10, 20, 30, 40, 50}
	lowCap := make([]int, cap(bigCap))            //Capacity 5
	copy(lowCap, []int{60, 70, 80, 90, 100, 200}) // 200 is on length 6 it would need 6 capacity to fill all on lowCap slice
	fmt.Println(lowCap)                           // [60 70 80 90 100] , 200 is lost since we have capacity for only 5 numbers

	// Better way to copy an array in my opinion

	letsCp := []int64{51, 52, 53, 54, 56, 57, 58}
	newCp := append([]int64(nil), letsCp...) // Create a nil slice of your type choice, remember that length is 0 and
	// This line above creates new backing arrays for every few things added , this because we know that capacity is 0
	// and each time capacity is surpassed it just doubles creating a new backing array

	newCp[4] = 5000
	fmt.Println(letsCp)
	fmt.Println(newCp) // Both copy methods are going to be creating a new backing array, the only way to
	// modify backing array from a slice, is creating a slice from a slice and then changing values within capacity

	//To avoid the problem that was mentioned above ,we can do the following
	//From the same letsCp slice

	perfoCp := append(make([]int64, 0, cap(letsCp)), letsCp...) // Make a slice from the same type of letsCp
	// start from length 0 (to copy from the beginning) to its full capacity (to append everything) and then append the wanted slice to be copied
	fmt.Println(perfoCp)

	// Inyect values into slice
	nums := []int{1, 3, 2, 4}
	nums = append(append(nums[0:2:2], 100, 200), nums[2:]...) // In place
	fmt.Println(nums)

	// Maybe more readable
	nums2 := []int{1, 3, 2, 4}
	temp := nums2[0:2:2]
	nums2 = append(append(temp, 100, 200), nums2[2:]...)
	fmt.Println(nums2)

	shapes := []string{"|", "/", "-", "\\", "|", "/", "-", "\\"}

	check := make([]int, 5, 10)
	check = append(check, 2)
	fmt.Println(check)
	fmt.Println(len(check))
	fmt.Println(cap(check))

	for x := 0; ; {
		fmt.Printf("\rPlease wait, processing... %s", shapes[x])
		time.Sleep(time.Millisecond * 150)
		x++
		if x == 8 {
			x = 0
		}
	}

}
