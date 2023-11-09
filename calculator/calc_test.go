package calc

import "testing"

// command
// go test . (run all file)
// go test ./... (run nested file from current dir)
// go test -v
func TestSumOfFirstThree(t *testing.T) {
	given := 3
	want := 6

	get := sumOfFirst(given)
	if want != get {
		t.Errorf("given %d want %d but get %d", given, want, get)
	}
}
