package os

import (
	"io/ioutil"
	"path"
	"strings"
)

// ReadDir scans the directory and return all environment variables, defined in its files
func ReadDir(dir string) (map[string]string, error) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(map[string]string)

	for _, info := range infos {
		if info.IsDir() {
			continue
		}

		p := path.Join(dir, info.Name())

		buf, err := ioutil.ReadFile(p)
		if err != nil {
			continue
		}

		content := string(buf)
		lines := strings.Split(content, "\n")

		env[info.Name()] = lines[0]
	}

	return env, nil
}

// RunCmd runs the command with the specified environment variables.
func RunCmd(cmd []string, env map[string]string) {

}
