package os

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

// initIO allow tests to setup standard IO streams for the command running via RunCmd function
func initIO(in io.Reader, out, err io.Writer) {
	stdin = in
	stdout = out
	stderr = err
}

// RunCmd runs the command with the specified environment variables.
func RunCmd(cmd []string, env map[string]string) {
	if len(cmd) < 1 {
		log.Println("Please input command to run")
		return
	}

	name := cmd[0]
	args := cmd[1:]

	envStrings := make([]string, len(env))
	idx := 0

	for name, value := range env {
		envStrings[idx] = fmt.Sprintf("%s=%s", name, value)
		idx++
	}

	command := exec.Command(name, args...)
	command.Env = append(command.Env, envStrings...)
	command.Stdin = stdin
	command.Stdout = stdout
	command.Stderr = stderr

	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
}
