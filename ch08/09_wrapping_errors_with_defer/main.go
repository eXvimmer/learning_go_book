package main

import "fmt"

func doThing1(val int) (string, error) {
	return "", nil
}

func doThing2(val string) (string, error) {
	return "", fmt.Errorf("doThing2: something bad happened")
}

func doThing3(val1, val2 string) (string, error) {
	return "", nil
}

func doSomethingBad(val1 int, val2 string) (string, error) {
	val3, err := doThing1(val1)
	if err != nil {
		return "", fmt.Errorf("in doSomethingBad: %w", err)
	}

	val4, err := doThing2(val2)
	if err != nil {
		return "", fmt.Errorf("in doSomethingBad: %w", err)
	}

	result, err := doThing3(val3, val4)
	if err != nil {
		return "", fmt.Errorf("in doSomethingBad: %w", err)
	}

	return result, nil
}

func doSomethingGood(val1 int, val2 string) (_ string, err error) {
	// NOTE: we hvae to name our return values so that we can refer to err in the
	// defferred function.
	defer func() {
		if err != nil {
			// wrapped error
			err = fmt.Errorf("in doSomethingGood: %w", err)
		}
	}()

	val3, err := doThing1(val1)
	if err != nil {
		return "", err
	}

	val4, err := doThing2(val2)
	if err != nil {
		return "", err
	}

	return doThing3(val3, val4)
}
