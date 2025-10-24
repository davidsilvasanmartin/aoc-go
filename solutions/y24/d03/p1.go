package d03

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type LastRead int

const (
	LastReadTrash = iota
	LastReadM
	LastReadU
	LastReadL
	LastReadOpeningParen
	LastReadFirstNum
	LastReadComma
	LastReadLastNum
)

func RunP1(input string) (string, error) {
	var state LastRead = LastReadTrash
	firstNumAsStr := ""
	lastNumAsStr := ""
	result := 0
	reader := bufio.NewReader(strings.NewReader(input))

	for {
		r, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}

		switch state {
		case LastReadTrash:
			if r == 'm' {
				state = LastReadM
			}
		case LastReadM:
			if r == 'm' {
				state = LastReadM
				firstNumAsStr = ""
				lastNumAsStr = ""
			} else if r == 'u' {
				state = LastReadU
			} else {
				resetP1State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case LastReadU:
			if r == 'm' {
				state = LastReadM
				firstNumAsStr = ""
				lastNumAsStr = ""
			} else if r == 'l' {
				state = LastReadL
			} else {
				resetP1State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case LastReadL:
			if r == 'm' {
				state = LastReadM
				firstNumAsStr = ""
				lastNumAsStr = ""
			} else if r == '(' {
				state = LastReadOpeningParen
			} else {
				resetP1State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case LastReadOpeningParen:
			if r == 'm' {
				state = LastReadM
				firstNumAsStr = ""
				lastNumAsStr = ""
			} else if '0' <= r && r <= '9' {
				firstNumAsStr += string(r)
				state = LastReadFirstNum
			} else {
				resetP1State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case LastReadFirstNum:
			if r == 'm' {
				state = LastReadM
				firstNumAsStr = ""
				lastNumAsStr = ""
			} else if '0' <= r && r <= '9' {
				firstNumAsStr += string(r)
			} else if r == ',' {
				state = LastReadComma
			} else {
				resetP1State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case LastReadComma:
			if r == 'm' {
				state = LastReadM
				firstNumAsStr = ""
				lastNumAsStr = ""
			} else if '0' <= r && r <= '9' {
				lastNumAsStr += string(r)
				state = LastReadLastNum
			} else {
				resetP1State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case LastReadLastNum:
			if r == 'm' {
				state = LastReadM
				firstNumAsStr = ""
				lastNumAsStr = ""
			} else if '0' <= r && r <= '9' {
				lastNumAsStr += string(r)
			} else if r == ')' {
				firstNum, _ := strconv.Atoi(firstNumAsStr)
				lastNum, _ := strconv.Atoi(lastNumAsStr)
				result += firstNum * lastNum
				resetP1State(&state, &firstNumAsStr, &lastNumAsStr)
			} else {
				resetP1State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		default:
			return "", errors.New("invalid state reached")
		}
	}

	return strconv.Itoa(result), nil
}

func resetP1State(state *LastRead, firstNumAsString *string, lastNumAsStr *string) {
	*state = LastReadTrash
	*firstNumAsString = ""
	*lastNumAsStr = ""
}
