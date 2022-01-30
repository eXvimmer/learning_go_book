package main

import (
	"fmt"
)

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder := Adder{start: 9}
	fmt.Println(myAdder.AddTo(4))

	f1 := myAdder.AddTo
	fmt.Println(f1(90))

	// NOTE: You can also create a function from the type itself. This is called
	// a method expression
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, -2))
}
