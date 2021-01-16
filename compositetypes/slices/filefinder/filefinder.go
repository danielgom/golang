package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Provide a directory please")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	ints := []int{1, 2, 3}
	news := append(make([]int, 0, len(ints)), ints...)
	fmt.Println(news)

	sort.Strings(args)
	var totBytes int
	for i := range files {
		if files[i].Size() == 0 {
			totBytes += len(files[i].Name()) + 2
		}
	}

	fmt.Println(totBytes)

	names := make([]byte, 0, totBytes)

	for i := range files {
		if files[i].Size() == 0 {
			names = strconv.AppendInt(names, int64(i) + 120, 10)
			names = append(names, files[i].Name()...) // Remember that a string can be added to byte slice since these are bytes
			names = append(names, '\n')
		}
	}


	err = ioutil.WriteFile("emptyFiles.txt", names, 0664)
	if err != nil {
		fmt.Println(err)
		return
	}
}
