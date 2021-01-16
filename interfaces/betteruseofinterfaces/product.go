package main

import (
	"fmt"
)

type product struct {
	Title    string
	Price    money
	Released timestamp
}

// Representation of the type (like a toString() of java), whenever printf is called then this method is called
func (p *product) String() string {
	return fmt.Sprintf("Title: %s, Price: %s, Released: %s", p.Title, p.Price, p.Released)
}

func (p *product) discount(ratio float64) {
	p.Price *= money(1 - ratio)
}

