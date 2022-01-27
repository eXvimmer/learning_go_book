package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// NOTE: we need to seed the generator
	// NOTE: n is scoped to the entire block of if else statemtn
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low.")
	} else if n > 5 {
		fmt.Println("That's too high.")
	} else {
		fmt.Println("That's a good number.")
	}
}
