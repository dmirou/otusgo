package main

import (
	"fmt"
	"os"

	ntpClient "github.com/dmirou/otus-go/hw-01-ntp-time/pkg/client"
)

func main() {
	const host = "0.beevik-ntp.pool.ntp.org"
	client := ntpClient.NewClient(host)
	timeNow, err := client.GetDate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't get time from %s, error: %s\n", host, err)
		os.Exit(1)
	}
	fmt.Printf("time now is %s\n", timeNow)
}
