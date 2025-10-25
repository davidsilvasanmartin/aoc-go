package d04

type Direction struct {
	r int
	c int
}

var (
	DirectionDown      *Direction = &Direction{r: 1, c: 0}
	DirectionUp        *Direction = &Direction{r: -1, c: 0}
	DirectionRight     *Direction = &Direction{r: 0, c: 1}
	DirectionLeft      *Direction = &Direction{r: 0, c: -1}
	DirectionDownRight *Direction = &Direction{r: 1, c: 1}
	DirectionUpRight   *Direction = &Direction{r: -1, c: 1}
	DirectionDownLeft  *Direction = &Direction{r: 1, c: -1}
	DirectionUpLeft    *Direction = &Direction{r: -1, c: -1}
)

var directions []*Direction = []*Direction{
	DirectionDown,
	DirectionUp,
	DirectionRight,
	DirectionLeft,
	DirectionDownRight,
	DirectionUpRight,
	DirectionDownLeft,
	DirectionUpLeft,
}

type TextGrid struct {
	grid   [][]rune
	target []rune
}

func (t *TextGrid) lastColIndex() int {
	return len((*t).grid[0]) - 1
}

func (t *TextGrid) lastRowIndex() int {
	return len((*t).grid) - 1
}

func (t *TextGrid) isDirectionSearchable(d Direction, fromRow int, fromCol int) bool {
	targetLenMinusOne := len(t.target) - 1
	isTargetFittingHorizontally := false
	isTargetFittingVertically := false

	if d.r >= 0 {
		isTargetFittingVertically = (fromRow + d.r*targetLenMinusOne) <= t.lastRowIndex()
	} else {
		isTargetFittingVertically = (fromRow + d.r*targetLenMinusOne) >= 0
	}
	if d.c >= 0 {
		isTargetFittingHorizontally = (fromCol + d.c*targetLenMinusOne) <= t.lastColIndex()
	} else {
		isTargetFittingHorizontally = (fromCol + d.c*targetLenMinusOne) >= 0
	}

	return isTargetFittingHorizontally && isTargetFittingVertically
}

func (t *TextGrid) isTargetFound(d Direction, fromRow int, fromCol int) bool {
	if !t.isDirectionSearchable(d, fromRow, fromCol) {
		return false
	}
	for i, r := range t.target {
		candidate := t.grid[fromRow+i*d.r][fromCol+i*d.c]
		if candidate != r {
			break
		} else if candidate == r && i == len(t.target)-1 {
			return true
		}
	}
	return false
}
