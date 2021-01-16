package main

type puzzle struct { // Embedding the product type brings its methods to the puzzle type so we are safe from removing them
	product
}

// Removing this method since product type has print and discount method already
/*
func (p *puzzle) print() {
	fmt.Printf("Title: %s, Price: %s\n", p.title, p.price.string())
}
*/
