package common

import (
	"slices"
	"testing"
)

func TestInputToIntSlices(t *testing.T) {
	input := " 1 22 \n   33  4"
	want := [][]int{{1, 22}, {33, 4}}

	got, err := InputToIntSlices(input)
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if !slices.Equal(got[0], want[0]) || !slices.Equal(got[1], want[1]) {
		t.Fatalf("Run(%q) = %q, want %q", input, got, want)
	}
}

func TestIsStrictlyIncreasing(t *testing.T) {
	tests := []struct {
		name     string
		first    int
		second   int
		expected bool
	}{
		{name: "increasing", first: -1, second: 3, expected: true},
		{name: "equal", first: 2, second: 2, expected: false},
		{name: "decreasing", first: 3, second: -1, expected: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsStrictlyIncreasing(tc.first, tc.second)
			if got != tc.expected {
				t.Errorf("first=%v, second=%v, got=%v, expected=%v", tc.first, tc.second, got, tc.expected)
			}
		})
	}
}

func TestIsPairSafe(t *testing.T) {
	tests := []struct {
		name     string
		first    int
		second   int
		inc      bool
		expected bool
	}{
		{name: "inc safe", first: 1, second: 2, inc: true, expected: true},
		{name: "inc safe 2", first: 1, second: 4, inc: true, expected: true},
		{name: "inc equal", first: 2, second: 2, inc: true, expected: false},
		{name: "inc not inc", first: 2, second: 1, inc: true, expected: false},
		{name: "inc over max jump", first: 1, second: 5, inc: true, expected: false},
		{name: "dec safe", first: 2, second: 1, inc: false, expected: true},
		{name: "dec safe 2", first: 4, second: 1, inc: false, expected: true},
		{name: "dec equal", first: 2, second: 2, inc: false, expected: false},
		{name: "dec not dec", first: 1, second: 2, inc: false, expected: false},
		{name: "inc over max jump", first: 5, second: 1, inc: false, expected: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPairSafe(tc.first, tc.second, tc.inc)
			if got != tc.expected {
				t.Errorf("first=%v, second=%v, inc=%v, got=%v, expected=%v", tc.first, tc.second, tc.inc, got, tc.expected)
			}
		})
	}
}
