package main

import (
	"fmt"
	"reflect"
)

type MulSects struct {
	FirstSection  Section `customTag:"first_section"`
	SecondSection Section
}

type Section struct {
	Title string
	Body  string
	Cta   string
}

type Foo struct {
	A int    `myTag:"value"`
	B string `myTag:"value2"`
}

func main() {

	m := &MulSects{
		FirstSection: Section{
			Title: "my first title",
			Body:  "my first body",
			Cta:   "my first cta",
		},
		SecondSection: Section{
			Title: "my second title",
			Body:  "my second body",
			Cta:   "my second cta",
		},
	}

	fmt.Println(m.FirstSection.Title)

	// Using reflection to dynamically access fields
	reflectValue := reflect.ValueOf(m).Elem()
	fmt.Println(m)
	fmt.Println(reflectValue)
	secVal := reflectValue.FieldByName("FirstSection")
	fmt.Println(reflect.TypeOf(*m).Field(0).Tag.Get("customTag")) // Get the tag of a field
	if secVal.IsValid() {
		fieldVal := secVal.FieldByName("Title")
		if fieldVal.IsValid() && fieldVal.CanSet() {
			fieldVal.SetString("Set value to the first title") // Set the value of the struct by reflection
		}
	}

	fmt.Println(m.FirstSection.Title) // Changed value with reflection

	var f Foo
	ft := reflect.TypeOf(f)
	fmt.Println(ft)
	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)
		fmt.Println(curField.Name, curField.Type.Name(),
			curField.Tag.Get("myTag")) // Check tags (metadata)
	}

}
