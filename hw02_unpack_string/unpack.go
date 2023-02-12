package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result strings.Builder

	prevChar := ""
	for _, char := range str {
		if unicode.IsDigit(char) {
			if prevChar == "" {
				return "", ErrInvalidString
			}
			result.WriteString(strings.Repeat(prevChar, int(char-'0')))
			prevChar = ""
		} else {
			result.WriteString(prevChar)
			prevChar = string(char)
		}
	}
	result.WriteString(prevChar)
	return result.String(), nil
}
