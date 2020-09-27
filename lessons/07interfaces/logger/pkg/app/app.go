package app

type Logger interface {
	Log(msg string) error
}

type App struct {
	logger Logger
}

func NewApp(l Logger) *App {
	return &App{
		logger: l,
	}
}

func (a *App) Run() {
	a.logger.Log("App started")
}
