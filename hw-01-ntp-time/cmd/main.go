package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	const host = "0.beevik-ntp.pool.ntp.org"
	response, err := ntp.Query(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't get time from %s, error: %s\n", host, err)
		os.Exit(1)
	}
	time := time.Now().Add(response.ClockOffset)
	fmt.Printf("time now is %s\n", time)
}
