package adder_test // NOTE: the package name

// NOTE: If you want to test just the public API of your package, Go has a
// convention for speci‐ fying this. You still keep your test source code in
// the same directory as the production source code, but you use
// packagename_test for the package name.

import (
	"testing"

	adder "github.com/exvimmer/testing_go/01_adder"
)

func TestAddNumbers(t *testing.T) {
	result := adder.AddNumber(4, 9)
	expected := 13
	if result != expected {
		t.Errorf("incorrect result: expected %d, got %d", expected, result)
	}
}
