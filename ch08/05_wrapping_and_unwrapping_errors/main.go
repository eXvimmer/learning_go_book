package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err) // wrappper
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		// NOTE: you don't usually call errors.Unwrap directly. Instead, you use
		// errors.Is and errors.As to find a specific wrapped error.
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr) // unwrapped error
		}
	}

	// NOTE: if you want to wrap an error with your custom error type, your error
	// type needs to implement the method Unwrap . This method takes in no
	// parameter and returns an error.
}
