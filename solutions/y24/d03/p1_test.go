package d03

import "testing"

func TestRunP1(t *testing.T) {
	input := ""
	want := ""

	got, err := RunP1(input)

	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if got != want {
		t.Fatalf("Run(%q) = %q, want %q", input, got, want)
	}
}
