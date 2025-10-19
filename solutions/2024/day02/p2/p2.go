package p2

import (
	"strconv"

	"github.com/davidsilvasanmartin/aoc-go/solutions/2024/day02/common"
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

	indexOfFirstElOfUnsafePair := getIndexOfFirstElOfUnsafePair(report)
	if indexOfFirstElOfUnsafePair == -1 {
		// If we have reached this point, it means that all pairs are safe.
		return true
	}

	// Try removing the element at index indexOfFirstElOfUnsafePair + 1 and perform a full check again
	// This is for cases like {7, 1, 8, 9} that would get fixed if we remove the second element of
	// the unsafe pair
	fixedReport := append(report[:indexOfFirstElOfUnsafePair], report[indexOfFirstElOfUnsafePair+1:]...)
	if len(fixedReport) < 2 {
		return true
	}
	isAnyPairUnsafe := getIndexOfFirstElOfUnsafePair(fixedReport) != -1
	if isAnyPairUnsafe {
		return false
	}

	// Finally, try removing the element at index indexOfFirstElOfUnsafePair and perform a full check again
	// This is for cases like {1, 7, 8, 9} that would get fixed if we remove the first element of
	// the unsafe pair
	fixedReport = append(report[:indexOfFirstElOfUnsafePair-1], report[indexOfFirstElOfUnsafePair:]...)
	if len(fixedReport) < 2 {
		return true
	}
	isAnyPairUnsafe = getIndexOfFirstElOfUnsafePair(fixedReport) != -1
	if isAnyPairUnsafe {
		return false
	}

	return true
}

// getIndexOfFirstUnsafePair returns the index of the first pair in the report that is deemed unsafe based on specific rules.
// If all pairs are safe, it returns -1.
func getIndexOfFirstElOfUnsafePair(report []int) int {
	first := report[0]
	second := report[1]
	inc := common.IsStrictlyIncreasing(first, second)

	for i := 0; i < len(report)-1; i++ {
		curr := report[i]
		next := report[i+1]
		if !common.IsPairSafe(curr, next, inc) {
			return i
		}
	}

	return -1
}
