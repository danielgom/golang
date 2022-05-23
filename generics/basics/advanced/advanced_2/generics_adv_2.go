package main

import "fmt"

type hero struct {
	name   string
	canFly bool
}

// Required method for the stringer interface
func (h *hero) String() string {
	return h.name
}

type joiner interface {
	join(string, fmt.Stringer) string
}

type commaJoiner struct{}

func (commaJoiner) join(agg string, s fmt.Stringer) string {

	// Just for demonstration, prefer strings.Builder to minimize memory copying
	if len(agg) > 0 {
		return agg + ", " + s.String()
	}
	return s.String()
}

// Type constraint

func join[S fmt.Stringer](items []S, j joiner) string {
	var result string

	for _, item := range items {
		// Item is a kind of stringer
		result = j.join(result, item)
	}
	return result
}

func main() {
	heroes := []*hero{
		{name: "Home lander", canFly: true},
		{name: "Starlight", canFly: false},
	}
	fmt.Print(join(heroes, commaJoiner{}))
}
