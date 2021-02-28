package main

import "fmt"

type metal struct {
	mass   float64
	volume float64
}

type gas struct {
	pressure      float64
	temperature   float64
	molecularMass float64
}

// Since we have a common behaviour in gas and metal structs then we can avoid creating different IsDenser methods

type Dense interface {
	Density() float64
}

// Density - return density of the gas
func (g *gas) Density() float64 {
	// density = ((molecularMass) * (pressure)) / ((gasConstant) * (Kelvin temperature))
	var density float64
	density = (g.molecularMass * g.pressure) / (0.0821) * (g.temperature + 273)
	return density
}

// Density - return density of the metal
func (m *metal) Density() float64 {
	return m.mass / m.volume
}

// String - this function implements string to format output
func (g *gas) String() string {
	return fmt.Sprintf("The pressure of the gas is %g, with temperature %g and Molecular Mass of %g\n ", g.pressure, g.temperature, g.molecularMass)
}

//IsDenser - compare density of object to other object
/*
func IsDenser(a, b *metal) bool {
	return a.Density() > b.Density()
}

*/

func IsDenser(a, b Dense) bool {
	return a.Density() > b.Density()
}

func main() {

	gold := metal{478, 24}
	silver := metal{100, 10}

	fmt.Println("Gold density", gold.Density())
	fmt.Println("Silver density", silver.Density())

	fmt.Println("Is gold denser than silver?", IsDenser(&gold, &silver))

	oxygen := gas{
		pressure:      5,
		temperature:   27,
		molecularMass: 32,
	}

	hydrogen := gas{
		pressure:      1,
		temperature:   0,
		molecularMass: 2,
	}

	fmt.Println("Is oxygen denser than hydrogen?", IsDenser(&oxygen, &hydrogen))
	fmt.Println("Is gold denser than oxygen?", IsDenser(&gold, &oxygen))

	fmt.Println(&oxygen)

}
