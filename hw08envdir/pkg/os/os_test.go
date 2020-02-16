package os

import (
	"reflect"
	"testing"
)

// TestReadDir check that the ReadDir function correctly
// reads environment vars from files in the specified dir
func TestReadDir(t *testing.T) {
	testData := []struct {
		dir     string
		willErr bool
		env     map[string]string
	}{
		{
			"./testenvdir",
			false,
			map[string]string{
				"VAR_1_UPPER":            "VALUE_1_UPPER",
				"var_2_lower":            "value_2_lower",
				"VAR_3-DASH":             "VALUE_3-DASH",
				"VAR_4_MULTILINE":        "VAR4_FIRST_LINE",
				"VAR_5_TWO_WORD_IN_LINE": "VALUE_5_FIRST_WORD VALUE_6_SECOND_WORD",
				"VAR_6_EQUAL_=":          "VALUE_6_EQUAL_=",
				"VAR_7.txt":              "VALUE_7_TXT",
			},
		},
	}

	for _, td := range testData {
		actual, err := ReadDir(td.dir)
		if td.willErr && err == nil {
			t.Fatalf("error is expected, but not received with dir: %q", td.dir)
			continue
		}

		if !td.willErr && err != nil {
			t.Fatalf("error is unexpected with dir: %q, err: %v", td.dir, err)
			continue
		}

		if !reflect.DeepEqual(td.env, actual) {
			t.Fatalf("env vars not equal with dir: %q,\nexpected: %q\nactual: %q",
				td.dir, td.env, actual)
			continue
		}
	}
}
