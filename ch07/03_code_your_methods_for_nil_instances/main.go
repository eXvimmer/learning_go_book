package main

import (
	"fmt"
)

// NOTE: What happens when you call a method on a nil instance: Go tries to
// invoke the method, if it's a method with a value receiver, you'll get a
// panic, as there is no value being pointed to by pointer. If it's a method
// with a pointer receiver, it can work if the methdo is written to handle the
// possibility of nil instance.

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(v int) *IntTree {
	if it == nil { // handling nil instance
		return &IntTree{val: v}
	}

	if v < it.val {
		it.left = it.left.Insert(v)
	} else if v > it.val {
		it.right = it.right.Insert(v)
	}

	return it
}

func (it *IntTree) Contains(v int) bool {
	switch {
	case it == nil: // handling nil instance
		return false
	case v < it.val:
		return it.left.Contains(v)
	case v > it.val:
		return it.right.Contains(v)
	default:
		return true
	}
}

func main() {
	var it *IntTree

	it = it.Insert(5) // Go allows you to call a method on a nil receiver
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))
}
