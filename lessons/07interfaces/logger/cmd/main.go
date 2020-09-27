package main

import (
	"os"

	"github.com/dmirou/otusgo/lessons/07interfaces/logger/pkg/app"
	"github.com/dmirou/otusgo/lessons/07interfaces/logger/pkg/logger"
)

func main() {
	a := app.NewApp(logger.NewLogger(1, os.Stdout))
	a.Run()
}
