package main

import (
	"errors"
	"fmt"
	"os"
)

func divMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is zero")
		// NOTE: error messages should'nt be capitalized nor should they end with
		// punctuation or a new line.
	}

	return numerator / denominator, numerator % denominator, nil
}

func main() {
	numerator, denominator := 20, 3
	remainder, mod, err := divMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder, mod)
}
