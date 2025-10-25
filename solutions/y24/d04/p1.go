package d04

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func (t *TextGrid) getTargetFoundCountP1() int {
	targetFoundCount := 0
	for row := 0; row <= t.lastRowIndex(); row++ {
		for col := 0; col <= t.lastColIndex(); col++ {
			for _, d := range directions {
				if t.isTargetFound(*d, row, col) {
					targetFoundCount++
				}
			}
		}
	}
	return targetFoundCount
}

func RunP1(input string) (string, error) {
	t := &TextGrid{
		grid:   [][]rune{},
		target: []rune{'X', 'M', 'A', 'S'},
	}
	currentRow := []rune{}
	reader := bufio.NewReader(strings.NewReader(input))

	for {
		r, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
		if r == '\n' {
			t.grid = append(t.grid, currentRow)
			currentRow = []rune{}
		} else {
			currentRow = append(currentRow, r)
		}
	}
	t.grid = append(t.grid, currentRow)

	return strconv.Itoa(t.getTargetFoundCountP1()), nil
}
