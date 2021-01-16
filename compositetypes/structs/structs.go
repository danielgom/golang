package main

import (
	"fmt"
	"github.com/dgomez/learngo/first/explain/compositetypes/structs/newlib"
	"reflect"
	"time"
)

type Person struct {
	Name   string
	job    string
	age    int
	height float64
}

type song struct {
	title  string
	artist string
}

type playlist struct {
	genre string
	songs []song
}

type weather struct {
	time   time.Time
	places map[string]playground
}

type playgrounds struct {
	city   string
	places map[string][]playground
}

type playground struct {
	number int
	name   string
}

func main() {

	daniel := Person{
		Name:   "Daniel",
		job:    "Developer",
		age:    24,
		height: 1.80,
	}

	katty := Person{
		Name:   "Catalina",
		job:    "Educator",
		age:    51,
		height: 1.64,
	}

	danielCopy := daniel

	fmt.Println(daniel)
	fmt.Printf("%#v\n", daniel) // Better way to print it
	fmt.Printf("%#v\n", katty)
	fmt.Printf("%#v\n", danielCopy)

	fmt.Println("Are these the same before changes? ", daniel == danielCopy) // this only compares if values within are the same , this is true
	// Change the copy
	danielCopy.age = 25
	danielCopy.job = "Data scientist"
	fmt.Printf("%#v\n", danielCopy) // Changed
	fmt.Printf("%#v\n", daniel)     // Not changed, very different from java , it seems that copy, creates another
	// struct type instead of pointing to the copied one

	fmt.Println("Are these the same after changes? ", daniel == danielCopy) // this is false, we changed danielCopy with lines 48 and 49

	rock := playlist{
		genre: "general rock",
		songs: []song{
			{"arks in the sky", "david bowie"},
			{"stairway to heaven", "zeppelin"},
			{"highway to hell", "ac/dc"},
		},
	}

	rock3 := playlist{
		genre: "general rock",
		songs: []song{
			{"arks in the sky", "david bowie"},
			{"stairway to heaven", "zeppelin"},
			{"highway to hell", "ac/dc"},
		},
	}

	fmt.Println(rock.songs[1].title)

	rock2 := rock

	//rock.songs = append(rock.songs, song{"this is new for me", "metallica"})
	//fmt.Println(rock)

	fmt.Println(rock2)
	fmt.Println(rock.genre == rock2.genre)

	for i := range rock.songs {
		rock.songs[i].title = "new songs" // This is going to change the tittle of every song
		// NOTE: IF WE DO NOT USE A VALUE IN THE FOR LOOP OVER SLICES THEN WE ARE AFFECTING THE ORIGINAL SLICE AND THE SLICES
		// THAT ARE POINTING TO THE SAME BACKING ARRAY
	}

	rock.genre = "new general rock"

	fmt.Println(rock)
	fmt.Println(rock2)
	fmt.Println("Comparison when a struct string is changed", rock.genre == rock2.genre)
	fmt.Println("Comparison when struct slice is changed", reflect.DeepEqual(rock.songs, rock2.songs))
	fmt.Println("This is rock   ", rock)
	fmt.Println("This is rock3   ", rock3)

	for _, v := range rock3.songs { // As we know with slices if we use value then it clones the slice
		v.title = "destroy" // EXAMPLE DIFFERENCE FROM THE FOR ABOVE
	}

	fmt.Println("Rock3 after for with values")
	fmt.Println(rock3) // same as when we declared rock3, no change from the for

	fmt.Println("Before student")

	obStudent := Student{ // Struct in the same package but in other file
		name:   "Daniel Gomez A.",
		age:    24,
		height: 1.80,
	}

	fmt.Println(obStudent)
	fmt.Println(reflect.TypeOf(obStudent))
	obStudent.age = 25
	fmt.Println(obStudent)

	// Testing struct with map and slices within

	climate := weather{
		time:   time.Now(),
		places: map[string]playground{
			"london": {number: 1, name: "England"},
			"mexico": {number: 2, name: "CDMX"}},
	}

	fmt.Println()
	fmt.Println(climate.time)

	climate2 := playgrounds{
		city: "Mexico City",
		places: map[string][]playground{
			"Veracruz": {{number: 1, name: "Xalapa"}, {number: 2, name: "Boca del rio"}},
		},
	}

	fmt.Println(climate2)

	// From another lib

	danielAcc := newlib.Account{ // We cannot create or call something from the other lib if that thing is not capitalized
		Name:    "Daniel Account", // To add values these fields must be public (not capitalized)
		Email:   "dga_355@hotmail.com",
		Balance: 56345.342,
	}

	fmt.Println(danielAcc)
	private := newlib.Private{}

	fmt.Println(private)
	fmt.Println(newlib.NewCellphone()) // If these values are not public , we can retrieve a new struct
	// by having the struct public and making a function like this one above
}
