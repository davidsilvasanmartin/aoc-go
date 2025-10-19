package math

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{name: "positive", input: 5, expected: 5},
		{name: "negative", input: -5, expected: 5},
		{name: "zero", input: 0, expected: 0},
		{name: "large positive", input: 123456789, expected: 123456789},
		{name: "large negative", input: -123456789, expected: 123456789},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Abs(tc.input)
			if got != tc.expected {
				t.Errorf("Abs(%d) = %d; want %d", tc.input, got, tc.expected)
			}
		})
	}
}
