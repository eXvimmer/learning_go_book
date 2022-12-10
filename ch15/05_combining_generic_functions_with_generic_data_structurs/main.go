package main

import (
	"fmt"
	"strings"
)

type BuiltInOrdered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~string | ~float64 | ~float32 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type OrderableFunc[T any] func(t1, t2 T) int

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}
	switch r := f(v, n.val); {
	case r <= -1:
		n.left = n.left.Add(f, v)
	case r >= 1:
		n.right = n.right.Add(f, v)
	}
	return n
}

func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}
	switch r := f(v, n.val); {
	case r <= -1:
		return n.left.Contains(f, v)
	case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
}

type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Order(other Person) int {
	out := strings.Compare(p.Name, other.Name)
	if out == 0 {
		out = p.Age - other.Age
	}
	return out
}

func OrderPeople(p1, p2 Person) int {
	out := strings.Compare(p1.Name, p2.Name)
	if out == 0 {
		out = p1.Age - p2.Age
	}
	return out
}

func BuiltInOrderable[T BuiltInOrdered](t1, t2 T) int {
	if t1 < t2 {
		return -1
	}
	if t1 > t2 {
		return 1
	}
	return 0
}

func main() {
	t1 := NewTree(BuiltInOrderable[int])
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)
	fmt.Println(t1.Contains(15))
	fmt.Println(t1.Contains(40))

	t2 := NewTree(OrderPeople)
	t2.Add(Person{"Mustafa", 30})
	t2.Add(Person{"Malena", 26})
	t2.Add(Person{"Emily", 25})
	fmt.Println(t2.Contains(Person{"Mustafa", 30}))
	fmt.Println(t2.Contains(Person{"Leslie", 30}))

	t3 := NewTree(Person.Order)
	t3.Add(Person{"Mustafa", 30})
	t3.Add(Person{"Malena", 26})
	t3.Add(Person{"Emily", 24})
	fmt.Println(t3.Contains(Person{"Mustafa", 30}))
	fmt.Println(t3.Contains(Person{"Leslie", 30}))
}
