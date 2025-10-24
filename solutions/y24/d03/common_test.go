package d03

import (
	"testing"
)

func TestSayHi(t *testing.T) {
	input := ""
	got, err := SayHi("")
	want := "Hi"

	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if got != want {
		t.Fatalf("Run(%q) = %q, want %q", input, got, want)
	}
}
