package slices

import "strconv"

func StringsToIntegers(strings []string) ([]int, error) {
	ints := make([]int, 0, len(strings))
	for _, s := range strings {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func RemoveAt(a []int, idx int) []int {
	b := make([]int, 0, len(a)-1)
	b = append(b, a[:idx]...)
	b = append(b, a[idx+1:]...)
	return b
}
