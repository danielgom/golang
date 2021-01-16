package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type game struct {
	Item  item   `json:"item"`
	Genre string `json:"genre"`
}

func main() {

	games := []game{
		{item{1, "god of war", 50}, "action adventure"},
		{item{2, "x-com 2", 30}, "strategy"},
		{item{3, "minecraft", 20}, "sandbox"},
	}

	fmt.Printf("%#v\n", games)
	marshal, err := json.MarshalIndent(games, "", "  ")
	if err != nil {
		return
	}

	fmt.Println(string(marshal))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > quit   : quits
`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for i := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					games[i].Item.Id, games[i].Item.Name, "("+games[i].Genre+")", games[i].Item.Price)
			}

		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}

			i, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			if len(games) <= i-1 {
				fmt.Println("sorry. i don't have the game")
				continue
			}

			fmt.Printf("#%d: %-15q %-20s $%d\n",
				games[i-1].Item.Id, games[i-1].Item.Name, "("+games[i-1].Genre+")", games[i-1].Item.Price)
		}
	}
}
