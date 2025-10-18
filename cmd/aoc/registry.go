package aoc

import (
	y2024d01p1 "github.com/davidsilvasanmartin/aoc-go/solutions/2024/day01/p1"
	y2024d01p2 "github.com/davidsilvasanmartin/aoc-go/solutions/2024/day01/p2"
)

var registry = map[int]map[int]map[int]RunnerFunc{
	2024: {
		1: {
			1: y2024d01p1.Run,
			2: y2024d01p2.Run,
		},
	},
}
