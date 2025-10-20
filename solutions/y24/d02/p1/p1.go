package p1

import (
	"strconv"

	"github.com/davidsilvasanmartin/aoc-go/solutions/y24/d02/common"
)

func Run(input string) (string, error) {
	reports, err := common.InputToIntSlices(input)
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

func isSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	first := report[0]
	second := report[1]
	inc := common.IsStrictlyIncreasing(first, second)

	for i := 0; i < len(report)-1; i++ {
		curr := report[i]
		next := report[i+1]
		if !common.IsPairSafe(curr, next, inc) {
			return false
		}
	}

	return true
}
