package d04

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func (t *TextGrid) getTargetFoundCountP2() int {
	targetFoundCount := 0
	// Ignore first and last indexes
	for row := 1; row < t.lastRowIndex(); row++ {
		for col := 1; col < t.lastColIndex(); col++ {
			// We are going to focus on finding the 'A' first
			charAtRowCol := t.grid[row][col]
			if charAtRowCol != 'A' {
				continue
			}
			// There should be no index out of bounds because we have ignored
			// the first and last column and row when looping
			rowUpIndex := row - 1
			rowDownIndex := row + 1
			colLeftIndex := col - 1
			colRightIndex := col + 1

			// Using the following functions is probably not performant, but
			// it reads well
			targetFoundBottomLeftToTopRight := t.isTargetFound(*DirectionUpRight, rowDownIndex, colLeftIndex)
			targetFoundTopRightToBottomLeft := t.isTargetFound(*DirectionDownLeft, rowUpIndex, colRightIndex)
			targetFoundTopLeftToBottomRight := t.isTargetFound(*DirectionDownRight, rowUpIndex, colLeftIndex)
			targetFoundBottomRightToTopLeft := t.isTargetFound(*DirectionUpLeft, rowDownIndex, colRightIndex)
			if (targetFoundBottomLeftToTopRight || targetFoundTopRightToBottomLeft) &&
				(targetFoundTopLeftToBottomRight || targetFoundBottomRightToTopLeft) {
				targetFoundCount++
			}
		}
	}
	return targetFoundCount
}

// RunP2 calculates the solution for the second part of the challenge.
// The code is very ugly and not generalizable, but it has been written
// like that on purpose to use the definitions and structures from the
// first part, to not write a lot of extra code
func RunP2(input string) (string, error) {
	t := &TextGrid{
		grid:   [][]rune{},
		target: []rune{'M', 'A', 'S'},
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

	return strconv.Itoa(t.getTargetFoundCountP2()), nil
}
