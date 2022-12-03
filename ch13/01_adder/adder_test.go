package adder

import "testing"

// NOTE: when testing unexported functions some people use an underscore
// between the word Test and the name of the function

func Test_addNumbers(t *testing.T) {
	result := addNumbers(4, 9)
	expected := 13
	if result != expected {
		t.Errorf("incorrect result: expected %d, got %d", expected, result)
	}

	// NOTE: If the failure of a check in a test means that further checks in the
	// same test function will always fail or cause the test to panic, use Fatal
	// or Fatalf. If you are testing several independent items (such as
	// validating fields in a struct), then use Error or Errorf so you can report
	// as many problems at once. This makes it easier to fix multiple prob‐ lems
	// without rerunning your tests over and over.
}
