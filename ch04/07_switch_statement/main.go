package main

import (
	"fmt"
)

func main() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	for _, word := range words {
		// NOTE: you can switch on any type that can be compared with ==
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
			// NOTE: By default, cases in switch statements in Go don't fall through
		case 5:
			wordlen := len(word) // block scoped (only visible here)
			fmt.Println(word, "is exactly the right length", wordlen)
		case 6, 7, 8, 9:
			// NOTE: in Go, an empty case means nothing happens
		default: // optional
			fmt.Println(word, "is a long word")
		}
		// NOTE: Go does include a keyword fallthrough, which lets one case to
		// continue on the next one; But think twice before using it.
	}
	fmt.Println("-----------------------------------------------------------")

	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop")
			break // NOTE: breaks out of the switch not the for loop
		default:
			fmt.Println(i, "is boring")
		}
	}
	// NOTE: to break out of the for loop we need to use a label
	fmt.Println("-----------------------------------------------------------")
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop")
			break loop // NOTE: exactly what we want
		default:
			fmt.Println(i, "is boring")
		}
	}
}
