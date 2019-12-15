package strings

import (
	"strconv"
	"strings"
	"unicode"
)

// Unpack replace sequence {char}{number} on sequence {char}{char}... with {number} length
// For example: a2b3cde -> aabbbcde
func Unpack(s string) (string, error) {
	var result strings.Builder
	var prev rune
	var prevNumber = false
	for _, r := range s {
		switch true {
		case !unicode.IsNumber(r):
			result.WriteRune(r)
			prev = r
			prevNumber = false
			continue
		case prevNumber || prev == 0:
			return "", &InvalidArgError{Value: s}
		}
		count, err := strconv.Atoi(string(r))
		if err != nil {
			return "", &InvalidArgError{Value: s}
		}
		result.WriteString(strings.Repeat(string(prev), count-1))
	}
	return result.String(), nil
}
