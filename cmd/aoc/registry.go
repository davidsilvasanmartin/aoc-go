package aoc

import (
	y24d01p1 "github.com/davidsilvasanmartin/aoc-go/solutions/y24/d01/p1"
	y24d01p2 "github.com/davidsilvasanmartin/aoc-go/solutions/y24/d01/p2"
	y24d02p1 "github.com/davidsilvasanmartin/aoc-go/solutions/y24/d02/p1"
	y24d02p2 "github.com/davidsilvasanmartin/aoc-go/solutions/y24/d02/p2"
	y24d03p1 "github.com/davidsilvasanmartin/aoc-go/solutions/y24/d03/p1"
	y24d03p2 "github.com/davidsilvasanmartin/aoc-go/solutions/y24/d03/p2"
)

var registry = map[int]map[int]map[int]RunnerFunc{
	2024: {
		1: {1: y24d01p1.Run, 2: y24d01p2.Run},
		2: {1: y24d02p1.Run, 2: y24d02p2.Run},
		3: {1: y24d03p1.Run, 2: y24d03p2.Run},
	},
}
