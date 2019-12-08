package main

import (
	"fmt"
	"os"

	"github.com/dmirou/otus-go/hw-01-ntp-time/pkg/client"
	"github.com/dmirou/otus-go/hw-01-ntp-time/pkg/client/transport/ntp"
)

func main() {
	const host = "0.beevik-ntp.pool.ntp.org"
	transport := ntp.NewTransport(host)
	client := client.NewClient(transport)
	timeNow, err := client.GetTime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't get time from %s, error: %s\n", host, err)
		os.Exit(1)
	}
	fmt.Printf("time now is %s\n", timeNow)
}
