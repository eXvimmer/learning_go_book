package main

import (
	"fmt"
)

func div(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}

func main() {
	result := div(5, 2)
	fmt.Println(result)
}
