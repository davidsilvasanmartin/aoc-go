package aoc

import (
	y24d01p1 "github.com/davidsilvasanmartin/aoc-go/solutions/2024/day01/p1"
	y24d01p2 "github.com/davidsilvasanmartin/aoc-go/solutions/2024/day01/p2"
	y24d02p1 "github.com/davidsilvasanmartin/aoc-go/solutions/2024/day02/p1"
	y24d02p2 "github.com/davidsilvasanmartin/aoc-go/solutions/2024/day02/p2"
)

var registry = map[int]map[int]map[int]RunnerFunc{
	2024: {
		1: {1: y24d01p1.Run, 2: y24d01p2.Run},
		2: {1: y24d02p1.Run, 2: y24d02p2.Run},
	},
}
