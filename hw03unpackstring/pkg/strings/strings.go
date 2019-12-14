package strings

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/dmirou/otusgo/hw03unpackstring/pkg/errors"
)

// Unpack replace sequence {char}{number} on sequence {char}{char}... with {number} length
// For example: a2b3cde -> aabbbcde
func Unpack(s string) (string, error) {
	var result strings.Builder
	var prev rune
	var prevNumber = false
	for _, r := range s {
		if !unicode.IsNumber(r) {
			result.WriteRune(r)
			prev = r
			prevNumber = false
			continue
		}
		if prevNumber || prev == 0 {
			return "", &errors.InvalidArgError{Value: s}
		}
		count, _ := strconv.Atoi(string(r))
		for i := 0; i < count-1; i++ {
			result.WriteRune(prev)
		}
	}
	return result.String(), nil
}
