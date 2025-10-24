package d02

import "testing"

func TestIsSafeP2(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{name: "empty", input: []int{}, expected: true},
		{name: "one el", input: []int{1}, expected: true},
		{name: "two equal", input: []int{1, 1}, expected: true},
		{name: "three equal", input: []int{1, 1, 1}, expected: false},
		{name: "four equal", input: []int{1, 1, 1, 1}, expected: false},
		{name: "inc safe 1", input: []int{1, 2, 3}, expected: true},
		{name: "inc safe 2", input: []int{1, 4, 7}, expected: true},
		{name: "dec safe 1", input: []int{3, 2, 1}, expected: true},
		{name: "dec safe 2", input: []int{7, 4, 1}, expected: true},
		{name: "inc jump over max", input: []int{1, 5}, expected: true},
		{name: "dec jump over max", input: []int{5, 1}, expected: true},
		{name: "dec then inc, safe if 2nd removed", input: []int{7, 0, 8, 9}, expected: true},
		{name: "equal then inc, safe if 1st or 2nd removed", input: []int{7, 7, 8, 9}, expected: true},
		{name: "inc then dec, safe if 2nd removed", input: []int{3, 9, 2, 1}, expected: true},
		{name: "equal then dec, safe if 1st or 2nd removed", input: []int{3, 3, 2, 1}, expected: true},
		{name: "inc safe if first removed", input: []int{1, 7, 8, 9}, expected: true},
		{name: "dec safe if first removed", input: []int{9, 3, 2, 1}, expected: true},
		{name: "inc safe if 3rd removed", input: []int{1, 2, 10, 3}, expected: true},
		{name: "inc safe if 3rd removed 2", input: []int{1, 2, 0, 3}, expected: true},
		{name: "inc safe if 3rd removed 3", input: []int{1, 2, 2, 3}, expected: true},
		{name: "dec safe if 3rd removed", input: []int{9, 8, 1, 7}, expected: true},
		{name: "dec safe if 3rd removed 2", input: []int{9, 8, 11, 7}, expected: true},
		{name: "dec safe if 3rd removed 2", input: []int{9, 8, 8, 7}, expected: true},
		{name: "inc safe if last removed", input: []int{1, 2, 3, 10}, expected: true},
		{name: "inc safe if last removed 2", input: []int{1, 2, 3, 2}, expected: true},
		{name: "inc safe if last removed 3", input: []int{1, 2, 3, 3}, expected: true},
		{name: "dec safe if last removed", input: []int{9, 8, 7, 1}, expected: true},
		{name: "dec safe if last removed 2", input: []int{9, 8, 7, 7}, expected: true},
		{name: "dec safe if last removed 3", input: []int{9, 8, 7, 8}, expected: true},
		{name: "inc dec with small jumps", input: []int{5, 7, 5, 4}, expected: true},
		{name: "dec inc with small jumps", input: []int{8, 6, 8, 9}, expected: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isSafeP2(tc.input)
			if !got == tc.expected {
				t.Errorf("input=%v, got=%v, expected=%v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestRunP2(t *testing.T) {
	input := "1 2 3\n1 1 1\n3 2 1"
	want := "2"

	got, err := RunP2(input)

	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if got != want {
		t.Fatalf("Run(%q) = %q, want %q", input, got, want)
	}
}
