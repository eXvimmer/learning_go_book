package main

import (
	"errors"
	"fmt"
	"os"
)

func divmod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return numerator / denominator, numerator % denominator, nil
}

func main() {
	div, remainder, err := divmod(13, 4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(div, remainder)
}
