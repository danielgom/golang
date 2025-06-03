package main

import (
	"github.com/pelletier/go-toml/v2"
	"log"
	"os"
	"strconv"
)

type Toml struct {
	General   map[string]any `toml:"general"`
	Cursor    map[string]any `toml:"cursor"`
	Font      map[string]any `toml:"font"`
	Keyboard  map[string]any `toml:"keyboard"`
	Selection map[string]any `toml:"selection"`
	Window    map[string]any `toml:"window"`
}

func main() {
	if len(os.Args) < 1 {
		log.Fatal("Please provide the font size")
	}

	file, err := os.ReadFile("/Users/daniel.gomez2/.config/alacritty/alacritty.toml")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Printing current file config: \n%s", string(file))

	var config Toml
	err = toml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}

	newFontSize, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatal(err)
	}

	config.Font["size"] = newFontSize

	newFile, err := toml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Printing new file config: \n%s", string(newFile))

	err = os.WriteFile("/Users/daniel.gomez2/.config/alacritty/alacritty.toml", newFile, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("File updated successfully")
}
