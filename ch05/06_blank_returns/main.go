package main

import (
	"errors"
	"fmt"
)

func divmod(numerator, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}

	result, remainder = numerator/denominator, numerator%denominator

	return // NOTE: naked return, this is confusing, don't use it
}

func main() {
	// NOTE: blank returns, never use these
	fmt.Println(divmod(13, 9))
}

