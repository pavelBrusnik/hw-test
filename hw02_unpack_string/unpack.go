package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	sb := strings.Builder{}
	runes, err := filterAndConvert(s)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(runes); i++ {
		currentRune := runes[i]
		var nextRune rune
		if unicode.IsDigit(currentRune) {
			continue
		}
		if i == len(runes)-1 {
			nextRune = '1'
		} else {
			nextRune = runes[i+1]
		}
		repeatCounter := 1
		if unicode.IsDigit(nextRune) {
			repeatCounter = int(nextRune - '0')
		}
		sb.WriteString(strings.Repeat(string(currentRune), repeatCounter))
	}
	return sb.String(), nil
}

func filterAndConvert(s string) ([]rune, error) {
	runes := []rune(s)
	if unicode.IsDigit(runes[0]) {
		return []rune{}, ErrInvalidString
	}
	for i := 0; i < len(runes)-1; i++ {
		currentRune := runes[i]
		nextRune := runes[i+1]
		if unicode.IsDigit(currentRune) && unicode.IsDigit(nextRune) {
			return []rune{}, ErrInvalidString
		}
	}

	return runes, nil
}
