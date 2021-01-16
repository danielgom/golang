package main

import (
	"encoding/json"
	"fmt"
)

type text struct {
	title string
	words int
}

type book struct {
	text text
	id   int64
}

type article struct {
	text   // Reduce amount of code
	author string
	title  string
}

type permissions map[string]bool

type user struct {
	Name     string `json:"name"`     // this is going to print the json name as name instead of Name
	Password string `json:"password"` // json field : password instead of Password
	//permissions //anonymous type
	Permissions permissions `json:"permissions,omitempty"` // same as above but its going to prevent encoding if this is empty (nil permission)
}

func main() {

	book1 := book{
		text: text{"BBB1", 3563221},
		id:   62342,
	}

	article1 := article{
		text:   text{"Medical science 1", 6321},
		author: "ITCS",
		title:  "Text2",
	}

	fmt.Println(book1)
	fmt.Println(article1)

	fmt.Println(article1.title)      // Gets into the main title
	fmt.Println(article1.text.title) // This to get into the text title

	// Json encode

	users := []user{
		{"Daniel Gomez A.", "pass", permissions{"admin": false, "read": true, "write": false}},
		{"Rodrigo Diaz de vivar", "pass2", permissions{"admin": true, "read": true, "write": true}},
		{"Servant castañeda", "nonsense", permissions{"admin": false, "read": false, "write": false}},
		{"Newguy", "pass3", permissions{}},
	}

	fmt.Println(users)

	out, err := json.MarshalIndent(users, "", "  ") // this just pretty-prints the marshal
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out)) // json only encodes public fields, we need to change the private into public (just change the first letter to upper case)

}
