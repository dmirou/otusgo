package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dmirou/otusgo/hw07copyfile/pkg/file"
)

func main() {
	from, _ := filepath.Abs("./pkg/file/file.go")
	to, _ := filepath.Abs("./pkg/file/file.go.new")
	if err := file.CopyWithProgress(from, to, 0, 0); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
