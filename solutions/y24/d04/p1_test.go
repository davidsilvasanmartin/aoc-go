package d04

import "testing"

func Test_WorksInAllDirections(t *testing.T) {
	g := &TextGrid{
		grid: [][]rune{
			{'X', 'X', 'X', 'X', 'X'},
			{'X', 'I', 'I', 'I', 'X'},
			{'X', 'I', 'H', 'I', 'X'},
			{'X', 'I', 'I', 'I', 'X'},
			{'X', 'X', 'X', 'X', 'X'},
		},
		target: []rune{'H', 'I'},
	}
	tests :=
		[]struct {
			d    *Direction
			r    int
			c    int
			want bool
		}{
			// From (2,2)
			{d: DirectionUpLeft, r: 2, c: 2, want: true},
			{d: DirectionUp, r: 2, c: 2, want: true},
			{d: DirectionUpRight, r: 2, c: 2, want: true},
			{d: DirectionRight, r: 2, c: 2, want: true},
			{d: DirectionDownRight, r: 2, c: 2, want: true},
			{d: DirectionDown, r: 2, c: 2, want: true},
			{d: DirectionDownLeft, r: 2, c: 2, want: true},
			{d: DirectionLeft, r: 2, c: 2, want: true},
			// From (0,0)
			{d: DirectionUpLeft, r: 0, c: 0, want: false},
			{d: DirectionUp, r: 0, c: 0, want: false},
			{d: DirectionUpRight, r: 0, c: 0, want: false},
			{d: DirectionRight, r: 0, c: 0, want: false},
			{d: DirectionDownRight, r: 0, c: 0, want: false},
			{d: DirectionDown, r: 0, c: 0, want: false},
			{d: DirectionDownLeft, r: 0, c: 0, want: false},
			{d: DirectionLeft, r: 0, c: 0, want: false},
			// From (1,1)
			{d: DirectionUpLeft, r: 1, c: 1, want: false},
			{d: DirectionUp, r: 1, c: 1, want: false},
			{d: DirectionUpRight, r: 1, c: 1, want: false},
			{d: DirectionRight, r: 1, c: 1, want: false},
			{d: DirectionDownRight, r: 1, c: 1, want: false},
			{d: DirectionDown, r: 1, c: 1, want: false},
			{d: DirectionDownLeft, r: 1, c: 1, want: false},
			{d: DirectionLeft, r: 1, c: 1, want: false},
		}

	for _, tc := range tests {
		t.Run("test", func(t *testing.T) {
			got := g.isTargetFound(*tc.d, tc.r, tc.c)
			if got != tc.want {
				t.Errorf("direction=%v, got=%v, want=%v", tc.d, got, tc.want)
			}
		})
	}

	got := g.getTargetFoundCountP1()
	want := 8
	if got != want {
		t.Fatalf("got=%v, want=%v", got, want)
	}
}

func Test_WorksAtGridExtremes(t *testing.T) {
	g := &TextGrid{
		grid: [][]rune{
			{'H', 'I', 'H', 'I', 'H'},
			{'I', 'I', 'I', 'I', 'I'},
			{'H', 'I', 'X', 'I', 'H'},
			{'I', 'I', 'I', 'I', 'I'},
			{'H', 'I', 'H', 'I', 'H'},
		},
		target: []rune{'H', 'I'},
	}
	got := g.getTargetFoundCountP1()
	want := 32
	if got != want {
		t.Fatalf("got=%v, want=%v", got, want)
	}
}
