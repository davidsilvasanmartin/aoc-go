package d03

import (
	"bufio"
	"fmt"
	"strings"
)

func RunP1(input string) (string, error) {
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanRunes)

	for s.Scan() {
		fmt.Println(s.Text())
	}

	return "", nil
}
