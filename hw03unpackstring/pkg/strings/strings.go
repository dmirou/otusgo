package strings

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/dmirou/otusgo/hw03unpackstring/pkg/errors"
)

func Unpack(s string) (string, error) {
	var result strings.Builder
	var prev rune
	var prevNumber = false
	for _, r := range s {
		isNumber := unicode.IsNumber(r)

		if !isNumber {
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
