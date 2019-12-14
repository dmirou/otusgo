package strings

import "errors"

func Unpack(s string) (string, error) {
	return s, errors.New("method is not implement")
}
