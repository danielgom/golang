package main

import (
	"fmt"
	"regexp"
)

func main() {

	rx := regexp.MustCompile("[^a-zA-Z0-9 ]")

	check := "This is &%$·$&()/(()===awesome"
	fmt.Println(rx.ReplaceAllString(check, ""))

	/*

	in := bufio.NewScanner(os.Stdin)

	var lines int

	for in.Scan() {
		lines ++
		fmt.Println("Scanned text: ", in.Text())
	}

	fmt.Println("There are " + strconv.FormatInt(int64(lines), 10) + " in this file")

	 */

	//words := make(map[string]bool)

	//in := bufio.NewScanner(os.Stdin)

	/*
	in.Split(bufio.ScanWords) scan words instead of lines

	for in.Scan() {
		word := strings.ToLower(in.Text())

		if len(word) > 2 {
			words[word] = true
		}
	}

	 */


}
