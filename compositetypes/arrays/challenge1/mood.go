package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	if x := os.Args; len(x) < 2 {
		fmt.Println("[your name]")
		return
	}

	emos := [...]string{" Feels happy ☺️",
		" Feels doubt😅",
		" Feels awesome😎", " Feels super happy🤗",
		" Feels nice😋", " Feels angelic😇", " Feels with glasses🤓"}
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Daniel" + emos[rand.Intn(len(emos))])
}
