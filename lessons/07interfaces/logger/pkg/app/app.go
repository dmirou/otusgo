package app

import "github.com/dmirou/otusgo/lessons/07interfaces/logger/pkg/logger"

type App struct {
	logger *logger.Logger
}

func NewApp(l *logger.Logger) *App {
	return &App{
		logger: l,
	}
}

func (a *App) Run() {
	a.logger.Write("App started")
}
