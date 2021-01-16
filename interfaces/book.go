package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	product
	published interface{}
}

// USE POINTER RECEIVERS WHEN WE ARE GOING TO CHANGE THE STATE , OR WHEN DATA IS TOO LARGE

func (b *book) print() { // Do no use methods this way, methods are meant to be used and change the state of the object
	//Only mistake here is not changing the state of the object
	b.product.print()
	p := formatSwitch(b.published)
	fmt.Printf("\t, (%v)", p)
}

func (b book) printBookValue() { // Remember, only use methods to change state, not to just print the struct
	fmt.Printf("Title: %s, Price: %s\n", b.title, b.price.string())
}

func (b *book) changeTitle(t string) { // If we do not add the * then we would need to return the complete b struct and save it
	b.title = t
}

func (b book) changeTitleValue(t string) book { // Need to return the book and reassign since we are not using pointers
	b.title = t
	return b
}
func format(v interface{}) string {
	// Zero value of an interface value is nil
	_, ok := v.(bool)
	if v == nil || ok {
		return "unknown"
	}
	var t int
	if v, ok := v.(int); ok {
		t = v
	}

	if v, ok := v.(string); ok {
		t, _ = strconv.Atoi(v)
	}

	return time.Unix(int64(t), 0).String()
}

func formatSwitch(v interface{}) string {
	var t int
	switch v := v.(type) {
	case int: // If the type is an int
		t = v
	case string: // If the type is an string
		t, _ = strconv.Atoi(v)
	default: // If the type is neither an int or a string
		return "unknown"
	}
	// Formatting date
	const layout = "01/2006"
	return time.Unix(int64(t),0).Format(layout)
}
