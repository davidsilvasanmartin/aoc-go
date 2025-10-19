package math

// Abs returns the absolute value of the given integer
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
