package main

import (
	"errors"
	"fmt"
)

func divmod(numerator int, denominator int) (result int, remainder int, err error) {
	// NOTE: result, remainder and err are predeclared. They are initialized to
	// their zero values.
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}

	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func main() {
	res, rem, err := divmod(13, 9)
	fmt.Println(res, rem, err)
}
