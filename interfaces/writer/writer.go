package main

import (
	"container/list"
	"fmt"
	"unicode/utf8"
)

type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	*bc += ByteCounter(len(p))
	return len(p), nil
}

func main() {

	var b ByteCounter
	fmt.Fprintf(&b, "Hello world")
	fmt.Println(b)

	l := list.New()
	fmt.Println(l.PushFront(33))
	fmt.Println(l.PushFront(55))
	fmt.Println(l.PushBack(22))
	fmt.Println(l.PushFront(66)) // 66, 55, 33 , 22

	fmt.Println()
	fmt.Println()
	fmt.Println()

	for e := l.Front(); e != nil; e = e.Next() {

		fmt.Println(e.Next())
		fmt.Println(e.Prev())
	}

	v := "œ@@@¢¢¢¢¢¢"
	fmt.Println(len(v))
	fmt.Println(utf8.RuneCountInString(v))
}
