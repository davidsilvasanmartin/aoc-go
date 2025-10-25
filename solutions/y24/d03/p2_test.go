package d03

import "testing"

func TestRunP2(t *testing.T) {
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
			{name: "don't", input: "don't()mul(5,5)", expected: "0"},
			{name: "do don't do don't", input: "do()mul(1,1)don't()mul(5,5)do()mul(1,1)don't()mul(5,5)", expected: "2"},
			{name: "with trash", input: "aieflkdo()?sa[0wumul(1,1)s02jdon't()sd0982mul(5,5);oaido()09h4mul(1,1)sdijhdon't()s[08mul(5,5)[p's", expected: "2"},
			// mul - mul
			{name: "mmul", input: "mmul(3,4)", expected: "12"},
			{name: "mumul", input: "mumul(3,4)", expected: "12"},
			{name: "mulmul", input: "mulmul(3,4)", expected: "12"},
			{name: "mul(mul", input: "mul(mul(3,4)", expected: "12"},
			{name: "mul(2mul", input: "mul(2mul(3,4)", expected: "12"},
			{name: "mul(2,mul", input: "mul(2,mul(3,4)", expected: "12"},
			{name: "mul(2,2mul", input: "mul(2,2mul(3,4)", expected: "12"},
			// do - mul
			{name: "dmul", input: "dmul(3,4)", expected: "12"},
			{name: "domul", input: "domul(3,4)", expected: "12"},
			{name: "do(mul", input: "do(mul(3,4)", expected: "12"},
			// don't - mul
			{name: "donmul", input: "donmul(3,4)", expected: "12"},
			{name: "don'mul", input: "don'mul(3,4)", expected: "12"},
			{name: "don'tmul", input: "don'tmul(3,4)", expected: "12"},
			{name: "don't(mul", input: "don't(mul(3,4)", expected: "12"},
			// mul - do [Adding don't() at the front so we can test the do()]
			{name: "mdo", input: "don't()mdo()mul(3,4)", expected: "12"},
			{name: "mudo", input: "don't()mudo()mul(3,4)", expected: "12"},
			{name: "muldo", input: "don't()muldo()mul(3,4)", expected: "12"},
			{name: "mul(do", input: "don't()mul(do()mul(3,4)", expected: "12"},
			{name: "mul(2do", input: "don't()mul(2do()mul(3,4)", expected: "12"},
			{name: "mul(2,do", input: "don't()mul(2,do()mul(3,4)", expected: "12"},
			{name: "mul(2,2do", input: "don't()mul(2,2do()mul(3,4)", expected: "12"},
			// do - do [Adding don't() at the front so we can test the do()]
			{name: "ddo", input: "don't()ddo()mul(3,4)", expected: "12"},
			{name: "dodo", input: "don't()dodo()mul(3,4)", expected: "12"},
			{name: "do(do", input: "don't()do(do()mul(3,4)", expected: "12"},
			// don't - do [Adding don't() at the front so we can test the do()]
			{name: "dondo", input: "don't()dondo()mul(3,4)", expected: "12"},
			{name: "don'do", input: "don't()don'do()mul(3,4)", expected: "12"},
			{name: "don'tdo", input: "don't()don'tdo()mul(3,4)", expected: "12"},
			{name: "don't(do", input: "don't()don't(do()mul(3,4)", expected: "12"},
			// mul - don't
			{name: "mdon't", input: "mdon't()mul(3,4)", expected: "0"},
			{name: "mudon't", input: "mudon't()mul(3,4)", expected: "0"},
			{name: "muldon't", input: "muldon't()mul(3,4)", expected: "0"},
			{name: "mul(don't", input: "mul(don't()mul(3,4)", expected: "0"},
			{name: "mul(2don't", input: "mul(2don't()mul(3,4)", expected: "0"},
			{name: "mul(2,don't", input: "mul(2,don't()mul(3,4)", expected: "0"},
			{name: "mul(2,2don't", input: "mul(2,2don't()mul(3,4)", expected: "0"},
			// do - don't
			{name: "ddon't", input: "ddon't()mul(3,4)", expected: "0"},
			{name: "dodon't", input: "dodon't()mul(3,4)", expected: "0"},
			{name: "do(don't", input: "do(don't()mul(3,4)", expected: "0"},
			// don't - don't
			{name: "dondon't", input: "dondon't()mul(3,4)", expected: "0"},
			{name: "don'don't", input: "don'don't()mul(3,4)", expected: "0"},
			{name: "don'tdon't", input: "don'tdon't()mul(3,4)", expected: "0"},
			{name: "don't(don't", input: "don't(don't()mul(3,4)", expected: "0"},
		}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := RunP2(tc.input)
			if err != nil {
				t.Fatalf("Run returned error: %v", err)
			}
			if got != tc.expected {
				t.Errorf("input=%v, got=%v, expected=%v", tc.input, got, tc.expected)
			}
		})
	}
}
