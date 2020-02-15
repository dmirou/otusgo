package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dmirou/otusgo/hw07copyfile/pkg/file"
)

var src string
var dest string
var offset int
var limit int

const helpText = `
Usage: cp [OPTION]... -src SOURCE -dest DEST
Copy SOURCE file to DEST file.

Options:
- src 		(mandatory)	path of file to copy
- dst 		(mandatory)	path of destination file
- limit 	(optional)	maximum bytes to copy
- offset 	(optional)	file to read from

Examples:

With minimum options
	cp -src /tmp/from.txt -dst /tmp/to.txt

With maximum options
	cp -src /tmp/from.txt -dst /tmp/to.txt -offset 10 -limit 5`

// nolint: gochecknoinits
func init() {
	flag.StringVar(&src, "src", "", "path of file to copy")
	flag.StringVar(&dest, "dst", "", "path of destination file")
	flag.IntVar(&offset, "offset", 0, "file to read from")
	flag.IntVar(&limit, "limit", 0, "maximum bytes to copy")
}

func main() {
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println(helpText)
		os.Exit(0)
	}

	if src == "" {
		fmt.Println("Source file is missing. Please specify it with -src option.")
		os.Exit(1)
	}

	if dest == "" {
		fmt.Println("Destination file is missing. Please specify it with -dst option.")
		os.Exit(1)
	}

	if err := file.CopyWithProgress(src, dest, int64(limit), int64(offset)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
