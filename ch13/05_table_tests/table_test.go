package table

import (
	"testing"
)

func TestDoMath(t *testing.T) {
	data := []struct {
		name       string
		num1, num2 int
		op         string
		expected   int
		errMsg     string
	}{
		{"addition", 4, 9, "+", 13, ""},
		{"subtraction", 9, 4, "-", 5, ""},
		{"multiplication", 4, 9, "*", 36, ""},
		{"division", 26, 2, "/", 13, ""},
		{"bad_division", 4, 0, "/", 0, "division by zero"},
		{"default_case", 4, 0, "@", 0, "unknown operator @"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
			// NOTE: Comparing error messages can be fragile, because there may not
			// be any compatibility guarantees on the message text. The function that
			// we are testing uses errors.New and fmt.Errorf to make errors, so the
			// only option is to compare the messages. If an error has a custom type,
			// use errors.Is or errors.As to check that the correct error is
			// returned.
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error message: `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
