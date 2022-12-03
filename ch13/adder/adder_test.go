package adder

import "testing"

// NOTE: when testing unexported functions some people use an underscore
// between the word Test and the name of the function

func Test_addNumbers(t *testing.T) {
	result := addNumbers(4, 9)
	expected := 13
	if result != expected {
		t.Error("incorrect result, expected", expected, "got", result)
	}
}
