package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		// anonymous function
		func(j int) {
			fmt.Println("Printing", j, "from inside of an anonymous function")
		}(i)
	}
}
