package d02

import (
	"strconv"

	"github.com/davidsilvasanmartin/aoc-go/internal/slices"
)

func RunP2(input string) (string, error) {
	reports, err := inputToIntSlices(input)
	if err != nil {
		return "", err
	}

	numOfSafeReports := 0

	for _, report := range reports {
		if isSafeP2(report) {
			numOfSafeReports++
		}
	}

	return strconv.Itoa(numOfSafeReports), nil
}

func isSafeP2(report []int) bool {
	if len(report) <= 2 {
		return true
	}

	indexOfFirstElOfUnsafePair := getIndexOfFirstElOfUnsafePair(report)
	isReportSafe := indexOfFirstElOfUnsafePair == -1
	if isReportSafe {
		return true
	}

	for i, _ := range report {
		candidate := slices.RemoveAt(report, i)
		isCandidateSafe := getIndexOfFirstElOfUnsafePair(candidate) == -1
		if isCandidateSafe {
			return true
		}
	}

	return false
}

// getIndexOfFirstElOfUnsafePair returns the index of the first pair in the report that is deemed unsafe based on specific rules.
// If all pairs are safe, it returns -1.
func getIndexOfFirstElOfUnsafePair(report []int) int {
	first := report[0]
	second := report[1]
	inc := isStrictlyIncreasing(first, second)

	for i := 0; i < len(report)-1; i++ {
		curr := report[i]
		next := report[i+1]
		if !isPairSafe(curr, next, inc) {
			return i
		}
	}

	return -1
}
