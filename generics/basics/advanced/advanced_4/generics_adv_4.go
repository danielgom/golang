package main

import "fmt"

type vectorElement interface {
	~int | ~uint | ~float32 | ~float64
}

type vector2d[T vectorElement] struct {
	X T
	Y T
}

func newVector2d[T vectorElement](x T, y T) vector2d[T] {
	return vector2d[T]{X: x, Y: y}
}

// add method can only add vectors of the same type
func (v1 vector2d[T]) add(v2 vector2d[T]) vector2d[T] {
	return vector2d[T]{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func main() {

	v1 := newVector2d(10.7, 20.1)
	v2 := newVector2d(12.3, 43.2)
	fmt.Println(v1.add(v2))
}
