package d01

import (
	"testing"
)

func TestRunP1(t *testing.T) {
	input := "1 2\n3 3\n20 10"
	want := "11"

	got, err := RunP1(input)

	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if got != want {
		t.Fatalf("Run(%q) = %q, want %q", input, got, want)
	}
}
