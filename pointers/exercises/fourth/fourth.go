package main

import "fmt"

func main() {
	schools := make([]map[int]string, 2) // Creating the slice with 2, remember that if we append then the first 2 will be empty
	for i := range schools {
		schools[i] = make(map[int]string)
	}

	load(schools[0], []string{"batman", "superman"})

	load(schools[1], []string{"spiderman", "wonder woman"})

	fmt.Println(schools[0])
	fmt.Println(schools[1])

	schools2 := make([]int, 0, 2)
	fmt.Println(cap(schools2))
	fmt.Println(len(schools2))
	schools2 = append(schools2, 3)
	fmt.Println(schools2)
	fmt.Println(schools2[:cap(schools2)])
}

func load(m map[int]string, students []string) {
	for i, name := range students {
		m[i+1] = name
	}
}
