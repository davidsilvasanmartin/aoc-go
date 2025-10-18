package p1

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Run(input string) (string, error) {
	return strconv.Itoa(len(input)), nil
}

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
