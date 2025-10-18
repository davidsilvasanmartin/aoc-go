package aoc

import y2024d01 "github.com/davidsilvasanmartin/aoc-go/solutions/2024/day01"

type runner func() error

var registry = map[int]map[int]runner{
	2024: {
		1: y2024d01.Run,
	},
}
