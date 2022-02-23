package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err) // wrapped error
	}
	f.Close()
	return nil
}

// NOTE: by default errors.Is uses == to compare each wrapped error with the
// specified error. If this doesn't work for an error type that you define,
// implement the Is method on your error.

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return reflect.DeepEqual(me, me2)
	}

	return false
}

// NOTE: Another use for defining your own Is method is to allow comparisons
// against errors that aren’t identical instances. You might want to pattern
// match your errors, specifying a filter instance that matches errors that
// have some of the same fields.

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}

func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code

		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}
	return false
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("The file doesn't exist")
		}
	}
}
