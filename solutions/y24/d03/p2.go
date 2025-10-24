package d03

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type P2LastRead int

const (
	P2LastReadTrash = iota
	P2LastReadD
	P2LastReadO
	P2LastReadDoOpeningParen
	P2LastReadDontN
	P2LastReadDontAp
	P2LastReadDontT
	P2LastReadDontOpeningParen
	P2LastReadMulM
	P2LastReadMulU
	P2LastReadMulL
	P2LastReadMulOpeningParen
	P2LastReadMulFirstNum
	P2LastReadMulComma
	P2LastReadMulLastNum
)

func RunP2(input string) (string, error) {
	var state P2LastRead = P2LastReadTrash
	firstNumAsStr := ""
	lastNumAsStr := ""
	result := 0
	isMulEnabled := true
	reader := bufio.NewReader(strings.NewReader(input))

	for {
		r, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}

		switch state {
		case P2LastReadTrash:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadMulM:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'u' {
				state = P2LastReadMulU
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadMulU:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'l' {
				state = P2LastReadMulL
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadMulL:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == '(' {
				state = P2LastReadMulOpeningParen
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadMulOpeningParen:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if '0' <= r && r <= '9' {
				firstNumAsStr += string(r)
				state = P2LastReadMulFirstNum
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadMulFirstNum:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if '0' <= r && r <= '9' {
				firstNumAsStr += string(r)
			} else if r == ',' {
				state = P2LastReadMulComma
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadMulComma:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if '0' <= r && r <= '9' {
				lastNumAsStr += string(r)
				state = P2LastReadMulLastNum
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadMulLastNum:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if '0' <= r && r <= '9' {
				lastNumAsStr += string(r)
			} else if r == ')' && isMulEnabled {
				firstNum, _ := strconv.Atoi(firstNumAsStr)
				lastNum, _ := strconv.Atoi(lastNumAsStr)
				result += firstNum * lastNum
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadD:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'o' {
				state = P2LastReadO
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadO:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == '(' {
				state = P2LastReadDoOpeningParen
			} else if r == 'n' {
				state = P2LastReadDontN
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadDoOpeningParen:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == ')' {
				isMulEnabled = true
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadDontN:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == '\'' {
				state = P2LastReadDontAp
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadDontAp:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 't' {
				state = P2LastReadDontT
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadDontT:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == '(' {
				state = P2LastReadDontOpeningParen
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		case P2LastReadDontOpeningParen:
			if r == 'm' {
				restartFromM(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == 'd' {
				restartFromD(&state, &firstNumAsStr, &lastNumAsStr)
			} else if r == ')' {
				isMulEnabled = false
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			} else {
				resetP2State(&state, &firstNumAsStr, &lastNumAsStr)
			}
		default:
			return "", errors.New("invalid state reached")
		}
	}

	return strconv.Itoa(result), nil
}

func resetP2State(state *P2LastRead, firstNumAsString *string, lastNumAsString *string) {
	*state = P2LastReadTrash
	*firstNumAsString = ""
	*lastNumAsString = ""
}

func restartFromM(state *P2LastRead, firstNumAsString *string, lastNumAsString *string) {
	*state = P2LastReadMulM
	*firstNumAsString = ""
	*lastNumAsString = ""
}

func restartFromD(state *P2LastRead, firstNumAsString *string, lastNumAsString *string) {
	*state = P2LastReadD
	*firstNumAsString = ""
	*lastNumAsString = ""
}
