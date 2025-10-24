package d01

import (
	"maps"
	"testing"
)

func TestBuildTimesSeenMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected map[int]int
	}{
		{name: "happy path", input: []int{1, 1, 2}, expected: map[int]int{1: 2, 2: 1}},
		{name: "empty", input: []int{}, expected: map[int]int{}},
		{name: "single element", input: []int{42}, expected: map[int]int{42: 1}},
		{name: "all equal", input: []int{3, 3, 3, 3}, expected: map[int]int{3: 4}},
		{name: "no duplicates", input: []int{1, 2, 3}, expected: map[int]int{1: 1, 2: 1, 3: 1}},
		{name: "mixed signs and zero", input: []int{-5, 0, 1}, expected: map[int]int{-5: 1, 0: 1, 1: 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := buildTimesSeenMap(tc.input)
			if !maps.Equal(got, tc.expected) {
				t.Errorf("buildTimesSeenMap(%v) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestRunP2(t *testing.T) {
	input := "1 10\n1 1\n100 1"
	want := "4"

	got, err := RunP2(input)

	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if got != want {
		t.Fatalf("Run(%q) = %q, want %q", input, got, want)
	}
}
