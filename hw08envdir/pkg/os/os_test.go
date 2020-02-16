package os

import (
	"bytes"
	"io/ioutil"
	"os"
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

type runCmdData struct {
	cmd []string
	env map[string]string
	out string
	err string
}

// TestRunCmd checks that RunCmd function runs specified command with specified environment variables
func TestRunCmd(t *testing.T) {
	testData := map[string]runCmdData{
		"one env var": {
			[]string{"printenv", "ENV_VAR"},
			map[string]string{
				"ENV_VAR": "test_env_var_value",
			},
			"test_env_var_value\n",
			"",
		},
		"two env var": {
			[]string{"printenv"},
			map[string]string{
				"ENV_VAR1": "test_env_var_value1",
				"ENV_VAR2": "test_env_var_value2",
			},
			"ENV_VAR1=test_env_var_value1\nENV_VAR2=test_env_var_value2\n",
			"",
		},
	}

	for name, td := range testData {
		testRunCmd(t, name, td)
	}
}

func testRunCmd(t *testing.T, name string, td runCmdData) {
	in := bytes.NewReader([]byte{})
	outBuf := bytes.NewBufferString("")
	errBuf := bytes.NewBufferString("")
	initIO(in, outBuf, errBuf)
	RunCmd(td.cmd, td.env)

	if outBuf.String() != td.out {
		t.Errorf("test %s, unexpected result in out stream: %q, expected: %q", name, outBuf.String(), td.out)
	}

	if errBuf.String() != td.err {
		t.Errorf("test %s, unexpected result in err stream: %q, expected: %q", name, errBuf.String(), td.err)
	}
}

// TestRunCmdWithStdErr checks that RunCmd function runs specified command with correct stdError stream
func TestRunCmdWithStdErr(t *testing.T) {
	content := `#!/bin/bash
echo "Hello from temp file" >&2`

	dir := "/tmp"

	tmpfile, err := ioutil.TempFile(dir, "example")
	if err != nil {
		t.Errorf("Can't create tmp file: %v", err)
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Errorf("Can't write file %q: %v", tmpfile.Name(), err)
	}

	if err := tmpfile.Close(); err != nil {
		t.Errorf("Can't close file  %q: %v", tmpfile.Name(), err)
	}

	testData := map[string]runCmdData{
		"echo to stderr": {
			[]string{tmpfile.Name()},
			map[string]string{},
			"",
			"Hello from temp file\n",
		},
	}

	for name, td := range testData {
		execFile := td.cmd[0]

		if err := os.Chmod(execFile, 0777); err != nil {
			t.Errorf("Can't change %s mode", execFile)
		}

		testRunCmd(t, name, td)
	}
}
