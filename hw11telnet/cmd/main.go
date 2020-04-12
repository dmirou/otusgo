package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dmirou/otusgo/hw11telnet/pkg/telnet"
)

const helpText = `Usage: gotelnet host port
The gotelnet command is used for interactive text communication with another
host (host:port) via TCP protocol.
Examples:
	gotelnet --timeout=10s host port
	gotelnet mysite.ru 8080
	gotelnet --timeout=3s 0.0.0.0 80`

func main() {
	var timeout string

	flag.StringVar(&timeout, "timeout", "10s", "connection timeout")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println(helpText)
		os.Exit(0)
	}

	if flag.NArg() < 2 {
		log.Fatal("Port is missing. Please specify it as a second argument.")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	fmt.Printf("host: %s, port: %s, timeout: %s\n", host, port, timeout)

	dur, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalf("Can't parse timeout: %v", err)
	}

	c := telnet.NewClient(host, port, dur)

	if err := c.Run(); err != nil {
		log.Fatalf("cannot run telnet client: %v", err)
	}
}
