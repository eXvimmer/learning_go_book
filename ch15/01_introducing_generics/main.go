package main

import "fmt"

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(v T) {
	s.vals = append(s.vals, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	back := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return back, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	v, ok := intStack.Pop()
	fmt.Println(v, ok) // 30, true

	fmt.Printf("contains 20: %t\n", intStack.Contains(20))
	fmt.Printf("contains 30: %t\n", intStack.Contains(30))
}
