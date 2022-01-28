package main

import (
	"fmt"
)

func main() {
	// NOTE: A regular switch only allows you to check a value for equality. A
	// blank switch allows you to use any boolean boolean comparison for each
	// case.

	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		// a blank switch
		switch wordLen := len(word); { // NOTE: you should end the statement with ;
		case wordLen < 5:
			fmt.Println(word, "is a short word")
		case wordLen > 10:
			fmt.Println(word, "is too long")
		default:
			fmt.Println(word, "is exactly the right length")
		}
	}
}
