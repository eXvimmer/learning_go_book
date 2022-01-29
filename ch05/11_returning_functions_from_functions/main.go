package main

import (
	"fmt"
)

func MakeMul(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	twoBase := MakeMul(2)
	threeBase := MakeMul(3)

	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}
