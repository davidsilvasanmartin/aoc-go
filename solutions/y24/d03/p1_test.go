package d03

import "testing"

func TestRunP1(t *testing.T) {
	tests :=
		[]struct {
			name     string
			input    string
			expected string
		}{
			{name: "single", input: "mul(3,4)", expected: "12"},
			{name: "start", input: "mul(3,4).moi?{}-204)", expected: "12"},
			{name: "none", input: "mul(3, 4).moi?{mul(3.4)}-204)mul{3,4)", expected: "0"},
			{name: "end", input: ".moi?{}-204)mul(3,4)", expected: "12"},
			{name: "several", input: "mul(3,4).moi,?{mul(3,4)}-204)mul(3,4)", expected: "36"},
			{name: "consecutive", input: "mul(3,4)mul(3,4)mul(3,4)", expected: "36"},
			{name: "mmul", input: "mmul(3,4)", expected: "12"},
			{name: "mumul", input: "mumul(3,4)", expected: "12"},
			{name: "mulmul", input: "mulmul(3,4)", expected: "12"},
			{name: "mul(mul", input: "mul(mul(3,4)", expected: "12"},
			{name: "mul(2mul", input: "mul(2mul(3,4)", expected: "12"},
			{name: "mul(2,mul", input: "mul(2,mul(3,4)", expected: "12"},
			{name: "mul(2,2mul", input: "mul(2,2mul(3,4)", expected: "12"},
		}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := RunP1(tc.input)
			if err != nil {
				t.Fatalf("Run returned error: %v", err)
			}
			if got != tc.expected {
				t.Errorf("input=%v, got=%v, expected=%v", tc.input, got, tc.expected)
			}
		})
	}
}
