package main

import (
	"fmt"
)

func main() {

	// When methods return more than 1 parameter these usually have names

	moby := book{title: "moby dick", price: 199.99}
	moby.printBook()
	moby.printBookValue()

	moby.changeTitle("this new name") // No need to reassign since we are using pointers
	moby.printBook()
	moby.printBookValue()

	moby = moby.changeTitleValue("this is second name") // Needed to reassign since we are not using pointers
	moby.printBook()

	cod := game{title: "Call of Duty", price: 99.99}
	cod.print()

	book.printBookValue(moby)

	var games []*game
	// Two ways of appending pointers to a slice of pointers
	games = append(games, []*game{{title: "minecraft", price: 19.99}, {title: "Cyberpunk", price: 199.99}} ...)
	games = append(games, &game{title: "MapleStory", price: 0}, &game{title: "League of legends", price: 0})
	// Print pointer slice
	fmt.Println("These are the adresses of the games", games)
	for i := range games {
		games[i].print()
	}

	fmt.Println()
	fmt.Println()
	gameList := list(games)
	gameList.print()

	fmt.Println()
	fmt.Println()

	var games2 []*game
	list(games2).print()

}
