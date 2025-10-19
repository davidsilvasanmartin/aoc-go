package p1

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/davidsilvasanmartin/aoc-go/internal/math"
)

func Run(input string) (string, error) {
	// Since inputToIntSlices checks that every line has exactly two integer numbers, we
	// won't check that the first and second slices have the same size
	firstSlice, secondSlice, err := inputToIntSlices(input)
	if err != nil {
		return "", err
	}

	slices.Sort(firstSlice)
	slices.Sort(secondSlice)

	difs := 0

	for index, firstNum := range firstSlice {
		secondNum := secondSlice[index]
		difs += math.Abs(firstNum - secondNum)
	}

	return strconv.Itoa(difs), nil
}

// inputToIntSlices checks the input line by line and reads the ints that are present
// in each line into two slices. It checks that every line of the input contains
// exactly two integer numbers
func inputToIntSlices(input string) ([]int, []int, error) {
	firstSlice := []int{}
	secondSlice := []int{}

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		line := sc.Text()
		nums := strings.Fields(line)
		if len(nums) != 2 {
			return nil, nil, fmt.Errorf("could not extract numbers from malformed string: %s", line)
		}

		firstNum, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, nil, fmt.Errorf("could not convert to int: %s", nums[0])
		}
		firstSlice = append(firstSlice, firstNum)

		secondNum, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, nil, fmt.Errorf("could not convert to int: %s", nums[1])
		}
		secondSlice = append(secondSlice, secondNum)
	}

	return firstSlice, secondSlice, nil
}
