package app

import (
	"fmt"

	"github.com/dmirou/otusgo/calendar/pkg/event"
	"github.com/dmirou/otusgo/calendar/pkg/event/repository/localcache"
	"github.com/dmirou/otusgo/calendar/pkg/event/usecase"
)

type App struct {
	euc event.UseCase
}

func New() *App {
	repo := localcache.New()
	uc := usecase.New(repo)

	return &App{euc: uc}
}

func (a *App) Run() {
	fmt.Println("Hello! I am calendar!")
}
