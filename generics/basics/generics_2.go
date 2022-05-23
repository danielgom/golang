package main

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// LinkedList Creating a struct (like a class with a type) any refers to any type
type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
}

type node[T any] struct {
	next  *node[T]
	value T
}

func (lst *LinkedList[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &node[T]{value: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &node[T]{value: v}
		lst.tail = lst.tail.next
	}
}

func (lst *LinkedList[T]) GetAll() []T {
	nodes := make([]T, 0, 0)
	for nd := lst.head; nd != nil; nd = nd.next {
		nodes = append(nodes, nd.value)
	}
	return nodes
}
