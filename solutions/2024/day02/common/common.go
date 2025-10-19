package common

import (
	"bufio"
	"strings"

	"github.com/davidsilvasanmartin/aoc-go/internal/slices"
)

func InputToIntSlices(input string) ([][]int, error) {
	intSlices := [][]int{}

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		line := sc.Text()
		numsStr := strings.Fields(line)
		numsInt, err := slices.StringsToIntegers(numsStr)
		if err != nil {
			return nil, err
		}
		intSlices = append(intSlices, numsInt)
	}

	return intSlices, nil
}

// IsStrictlyIncreasing returns whether the sequence (first, second) is strictly increasing or not.
// Returns false if first == second
func IsStrictlyIncreasing(first int, second int) bool {
	return (second - first) > 0
}
