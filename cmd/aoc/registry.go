package aoc

import (
	y24d01 "github.com/davidsilvasanmartin/aoc-go/solutions/y24/d01"
	y24d02 "github.com/davidsilvasanmartin/aoc-go/solutions/y24/d02"
	y24d03 "github.com/davidsilvasanmartin/aoc-go/solutions/y24/d03"
	y24d04 "github.com/davidsilvasanmartin/aoc-go/solutions/y24/d04"
)

var registry = map[int]map[int]map[int]RunnerFunc{
	2024: {
		1: {1: y24d01.RunP1, 2: y24d01.RunP2},
		2: {1: y24d02.RunP1, 2: y24d02.RunP2},
		3: {1: y24d03.RunP1, 2: y24d03.RunP2},
		4: {1: y24d04.RunP1, 2: y24d04.RunP2},
	},
}
