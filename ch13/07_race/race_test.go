package race

import "testing"

func TestGetCounter(t *testing.T) {
	got := getCounter()
	want := 5000
	if want != got {
		t.Errorf("wanted: %d, got: %d", want, got)
	}
}
