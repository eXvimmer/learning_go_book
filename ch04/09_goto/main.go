package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// NOTE: In Go, a goto statement specifies a labeled line of code and
	// execution jumps to it. However, you can't jump anywhere. Go forbids jumps
	// that skip over variable declarations and jumps that go into an inner or
	// parallel block.

	a := rand.Intn(10) // not seeded
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	// We don't want to run the next line if loop doesn't complete normally
	fmt.Println("Do something when the loop completes normally")
done:
	fmt.Println("Do complicated stuff no matter why we left the loop")
	fmt.Println(a)
}
