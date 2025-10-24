package d01

import (
	"slices"
	"strconv"

	"github.com/davidsilvasanmartin/aoc-go/internal/math"
)

func RunP1(input string) (string, error) {
	// Since inputToIntSlices checks that every line has exactly two integer numbers, we
	// won't check that the first and second slices have the same size
	leftList, rightList, err := InputToIntSlices(input)
	if err != nil {
		return "", err
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	diffs := 0

	for index, firstNum := range leftList {
		secondNum := rightList[index]
		diffs += math.Abs(firstNum - secondNum)
	}

	return strconv.Itoa(diffs), nil
}
