package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {

	fmt.Println("Slices tutorial")
	// Small differences between slices and arrays
	var books [5]string
	books[0] = "dracula"
	books[1] = "slayer"
	books[2] = "angel"

	newBooks := [5]string{"dantes", "goood"}
	fmt.Println(books == newBooks) // the only way to compare is if newBooks is [5] stringtype

	games := []string{"Pokemon", "sims"}
	newGames := []string{"pacman", "maplestory", "doom"}
	fmt.Printf("books:     :%#v\n", books)
	fmt.Printf("newBooks:     :%#v\n", newBooks)
	fmt.Printf("games:     :%#v\n", games)
	fmt.Printf("newGames:     :%#v\n", newGames) // Notice that slice type never changes (arrays are fixed length [4][3] etc etc)
	// slices are just []string []int etc etc , not fixed length

	newGames = games // Declaring slices from other slices is easier,
	// KEEP IN MIND THAT SOMETHING HAPPENS HERE, IF WE CHANGE A VALUE EITHER IN newGames[0] THEN THAT CHANGE IS GOING TO APPLY
	// IN games[0] VALUE, AND VICEVERSA , THE ONLY WAY TO NOT HAVE THIS IS IF ONE OF THEM GOES TO NIL
	games = nil // LIKE HERE
	fmt.Println(games)
	fmt.Println(newGames)

	// slices with nil are not initialized, slices with no values are initialized with no values
	var test1 []int
	var test2 []int
	test2 = nil
	fmt.Println(test1)
	fmt.Println(len(test1))
	fmt.Println(test2)
	fmt.Println(len(test2))

	if len(games) != len(test2) { // best way to compare slices since an empty slice and a nil slice value is 0 always
		fmt.Println("not equal")
	} else {
		fmt.Println("equal")
	}

	rand.Seed(time.Now().UnixNano())

	//maxVal, _ := strconv.Atoi(os.Args[1])
	maxVal := 22

	var uniques []int

loop:
	for len(uniques) < maxVal {
		n := rand.Intn(maxVal*maxVal) + 1
		for _, val := range uniques {
			if val == n {
				continue loop
			}
		}
		uniques = append(uniques, n)
	}

	fmt.Println("\n\nuniques:", uniques)
	sort.Ints(uniques)
	fmt.Println("\nsorted:", uniques)

	nums := [5]int{5, 4, 3, 2, 1}
	sort.Ints(nums[:]) // Sort can only apply in slices, so it needs to be converted to a slice
	fmt.Println(nums)

	/*
		fmt.Println(games)
		fmt.Println(append(games, "batman"))
		fmt.Println(games)
		fmt.Println(len(games))
	*/

	// Append

	first := []string{"dunno", "what", "is", "this"}
	second := []string{"maybe", "this", "is", "it"}

	first = append(first, "intermediate", "or", " advanced")
	first = append(first, "this is a single value")
	first = append(first, second...) // second... means that it is going to append all values
	numberOf := 12
	fmt.Println(first[numberOf-1])

	fmt.Println(first)
	fmt.Println(len(first))

	// Sorting

	un := []int{3, 12, 5, 2, 6, 234, 6, 23, 5121, 3, 56, 673}
	sort.Sort(sort.IntSlice(un)) // In place sorting, no need to reassign variables
	sort.Ints(un)                // Better way to sort
	fmt.Println(un)

	ints := append([]int8{}, 2, 3, 4, 5, 6, 7, 8)
	intsArr := []int64{1, 2, 3, 4, 5}
	strArr := []string{"a", "b", "c", "d"}

	weird := append([]byte("hello"), "world"...) // Same as before , all values, since it is a byte, every letter in
	// world is going to be converted to byte that's why we need to use "world"...
	fmt.Println(&weird[0])
	fmt.Println(&weird[1])
	fmt.Println(&ints[0])
	fmt.Println(&ints[1])
	fmt.Println(&intsArr[0])
	fmt.Println(&intsArr[1])
	fmt.Println(&intsArr[2])
	fmt.Println(&strArr[0])
	fmt.Println(&strArr[1])
	fmt.Println(&strArr[2])

	// Slicing
	items := []string{"pacman", "mario", " tetris", "doom", "galaga", "frogger",
		"asteroids", " simcity", " metroid", "defender", "rayman", "tempset", "ultima"}
	fmt.Println(items)
	top3 := items[5:12]
	fmt.Println(top3)
	fmt.Println(items)
	anotherSlice := top3[:3]
	fmt.Println(anotherSlice)

	fmt.Println(append(top3, "!"))
	fmt.Println(items)
	const pageSize = 4
	for x := 0; x < len(items); x += pageSize {
		to := x + pageSize
		if to > len(items) {
			to = len(items)
		}
		page := items[x:to]
		head := fmt.Sprintf("Page #%d", (x/pageSize)+1) // Sprintf returns the formatted string instead of printing it fast
		fmt.Println(head, page)
	}

	nextT := "Hello how are you"
	fmt.Println(strings.Fields(nextT))

	lines := "trimming spaces ÷÷∞¢#¢©√ß∫∫®€æ∑å∫∂œæ€Ωå∫∂  "
	fmt.Println(lines + "nope")
	fmt.Println(strings.TrimSpace(lines) + "nope")

	fmt.Println(len(lines))
	fmt.Println(utf8.RuneCountInString(lines))

	check1 := make([]int, 0, 10)
	fmt.Println(cap(check1))
	fmt.Println(len(check1))
	fmt.Println(check1[0:cap(check1)])
	fundModSl(&check1)
	fmt.Println(check1)

}

func fundModSl(sl *[]int) {
	*sl = append(*sl, 10, 20, 30)
}
