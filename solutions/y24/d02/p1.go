package d02

import (
	"strconv"
)

func RunP1(input string) (string, error) {
	reports, err := InputToIntSlices(input)
	if err != nil {
		return "", err
	}

	numOfSafeReports := 0

	for _, report := range reports {
		if isSafeP1(report) {
			numOfSafeReports++
		}
	}

	return strconv.Itoa(numOfSafeReports), nil
}

func isSafeP1(report []int) bool {
	if len(report) < 2 {
		return true
	}

	first := report[0]
	second := report[1]
	inc := IsStrictlyIncreasing(first, second)

	for i := 0; i < len(report)-1; i++ {
		curr := report[i]
		next := report[i+1]
		if !IsPairSafe(curr, next, inc) {
			return false
		}
	}

	return true
}
