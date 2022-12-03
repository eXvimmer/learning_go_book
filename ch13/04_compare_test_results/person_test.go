package person

// NOTE: you need to install github.com/google/go-cmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})
	expected := Person{
		Name: "Mustafa",
		Age:  30,
	}
	result := createPerson("Mustafa", 30)
	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
}
