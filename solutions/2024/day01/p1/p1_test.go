package p1

import (
	"slices"
	"testing"
)

func TestRun(t *testing.T) {
	input := "Hi"
	want := "2"

	got, err := Run(input)

	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if got != want {
		t.Fatalf("Run(%q) = %q, want %q", input, got, want)
	}
}

func TestInputToIntSlices(t *testing.T) {
	input := " 1 22 "
	wantFirstSlice := []int{1}
	wantSecondSlice := []int{22}

	gotFirstSlice, gotSecondSlice, err := inputToIntSlices(input)
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if !slices.Equal(gotFirstSlice, wantFirstSlice) {
		t.Fatalf("Run(%q) = %q, want %q", input, gotFirstSlice, wantFirstSlice)
	}
	if !slices.Equal(gotSecondSlice, wantSecondSlice) {
		t.Fatalf("Run(%q) = %q, want %q", input, gotSecondSlice, wantSecondSlice)
	}
}
