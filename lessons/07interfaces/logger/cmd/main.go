package main

import (
	"os"

	"github.com/dmirou/otusgo/lessons/07interfaces/logger/pkg/logger"
)

func main() {
	log := logger.NewLogger(1, os.Stdout)
	log.Write("Hello!")
}
