package cleanup

import (
	"os"
	"testing"
)

// createFile is a helper function called from multiple tests
func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}

	// do something with f ...

	// run when the test and all its subsets are complete
	t.Cleanup(func() {
		os.Remove(f.Name())
	})

	return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
	_, err := createFile(t)
	if err != nil {
		t.Fatalf("failed: %v", err)
	}

	// do testing, don't worry about cleanup
}
