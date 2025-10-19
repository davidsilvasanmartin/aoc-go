package p1

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/davidsilvasanmartin/aoc-go/internal/math"
	"github.com/davidsilvasanmartin/aoc-go/internal/slices"
)

func Run(input string) (string, error) {
	reports, err := inputToIntSlices(input)
	if err != nil {
		return "", err
	}

	numOfSafeReports := 0

	for _, report := range reports {
		if isSafe(report) {
			numOfSafeReports++
		}
	}

	return strconv.Itoa(numOfSafeReports), nil
}

func inputToIntSlices(input string) ([][]int, error) {
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

func isSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	first := report[0]
	second := report[1]
	inc := isInc(first, second)

	for i := 0; i < len(report)-1; i++ {
		curr := report[i]
		next := report[i+1]
		if isInc(curr, next) != inc {
			return false
		}
		diff := math.Abs(next - curr)
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func isInc(first int, second int) bool {
	return (second - first) > 0
}
