package main

type game struct { // Type game will have the methods of the product type thats why we are going to comment them
	product
}

// Since the product type has these methods already, then we are save from removing them

/*
func (g *game) print() {
	fmt.Printf("Title: %s, Price: %s\n", g.title, g.price.string())
}

func (g *game) discount(ratio float64) {
	g.price *= money(1 - ratio)
}
*/
