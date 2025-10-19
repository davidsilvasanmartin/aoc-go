package p1

import "testing"

func TestIsSafe(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{name: "empty", input: []int{}, expected: true},
		{name: "one el", input: []int{1}, expected: true},
		{name: "two equal", input: []int{1, 1}, expected: false},
		{name: "three equal", input: []int{1, 1, 1}, expected: false},
		{name: "four equal", input: []int{1, 1, 1, 1}, expected: false},
		{name: "inc two equal last", input: []int{1, 2, 3, 3}, expected: false},
		{name: "dec two equal last", input: []int{3, 2, 1, 1}, expected: false},
		{name: "inc safe 1", input: []int{1, 2, 3}, expected: true},
		{name: "inc safe 2", input: []int{1, 4, 7}, expected: true},
		{name: "dec safe 1", input: []int{3, 2, 1}, expected: true},
		{name: "dec safe 2", input: []int{7, 4, 1}, expected: true},
		{name: "inc jump over max", input: []int{1, 5}, expected: false},
		{name: "inc jump over max at end", input: []int{1, 2, 3, 7}, expected: false},
		{name: "dec jump over max", input: []int{5, 1}, expected: false},
		{name: "dec jump over max at end", input: []int{7, 6, 5, 1}, expected: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isSafe(tc.input)
			if !got == tc.expected {
				t.Errorf("input=%v, got=%v, expected=%v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestRun(t *testing.T) {
	input := "1 2\n1 1\n2 1"
	want := "2"

	got, err := Run(input)

	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if got != want {
		t.Fatalf("Run(%q) = %q, want %q", input, got, want)
	}
}
