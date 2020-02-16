package main

import (
	"fmt"
	"log"
	"os"

	dos "github.com/dmirou/otusgo/hw08envdir/pkg/os"
)

const helpText = `Usage: goenv /path/to/evndir command arg1 arg2 ...
Load environment variables from /path/to/evndir and run command with args and the result env vars.

Examples:
	goenv /tmp/envdir echo $VAR1
	goenv ./pkg/os/testenvdir/ printenv`

func main() {
	if len(os.Args) == 1 {
		fmt.Println(helpText)
		os.Exit(0)
	}

	if len(os.Args) == 2 {
		log.Fatal("Command is missing. Please specify it as a second argument.")
	}

	vars, err := dos.ReadDir(os.Args[1])
	if err != nil {
		log.Fatalf("Can't load env from directory %q: %v", os.Args[1], err)
	}

	// skip current binary path and directory path
	cmd := os.Args[2:]
	dos.RunCmd(cmd, vars)
}
