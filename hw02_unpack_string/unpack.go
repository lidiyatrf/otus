package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	sb := strings.Builder{}
	var lastSymbol rune
	var escaped bool

	for _, s := range str {
		switch {
		case s == '\\':
			if lastSymbol != 0 {
				sb.WriteRune(lastSymbol)
			}
			if escaped {
				lastSymbol = s
			} else {
				lastSymbol = 0
			}
			escaped = !escaped
		case unicode.IsLower(s):
			if escaped {
				return "", ErrInvalidString
			}
			if lastSymbol != 0 {
				sb.WriteRune(lastSymbol)
			}
			lastSymbol = s
		case unicode.IsDigit(s):
			if escaped {
				escaped = false
				lastSymbol = s
				continue
			}
			if lastSymbol == 0 {
				return "", ErrInvalidString
			}
			if amount := int(s - '0'); amount > 0 {
				sb.WriteString(strings.Repeat(string(lastSymbol), amount))
			}
			lastSymbol = 0
		default:
			return "", ErrInvalidString
		}
	}
	if escaped {
		return "", ErrInvalidString
	}
	if lastSymbol != 0 {
		sb.WriteRune(lastSymbol)
	}

	return sb.String(), nil
}
