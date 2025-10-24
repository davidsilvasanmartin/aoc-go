package d01

import (
	"strconv"
)

func RunP2(input string) (string, error) {
	leftList, rightList, err := InputToIntSlices(input)
	if err != nil {
		return "", err
	}

	timesSeenMap := buildTimesSeenMap(rightList)
	simScore := 0

	for _, num := range leftList {
		timesSeen, ok := timesSeenMap[num]
		if ok {
			simScore += num * timesSeen
		}
	}

	return strconv.Itoa(simScore), nil
}

// buildTimesSeenMap creates and returns a map where keys are integers
// from the input slice and values are their counts
func buildTimesSeenMap(slice []int) map[int]int {
	timesSeenMap := map[int]int{}
	for _, num := range slice {
		timesSeenMap[num] += 1
	}
	return timesSeenMap
}
