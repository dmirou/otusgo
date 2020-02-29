package app

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dmirou/otusgo/calendar/pkg/config"
	"github.com/dmirou/otusgo/calendar/pkg/event"
	"github.com/dmirou/otusgo/calendar/pkg/event/repository/localcache"
	"github.com/dmirou/otusgo/calendar/pkg/event/usecase"
	"github.com/dmirou/otusgo/calendar/pkg/version"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type App struct {
	euc    event.UseCase
	cfg    *config.Config
	logger *zap.Logger
}

func New(configPath string) (*App, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("unexpected error during reading config file: %v", err)
	}

	cfg := config.Config{}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct, %v", err)
	}

	repo := localcache.New()
	uc := usecase.New(repo)

	l, err := makeLogger(&cfg)
	if err != nil {
		return nil, err
	}

	defer l.Sync() // nolint: errcheck

	return &App{euc: uc, cfg: &cfg, logger: l}, nil
}

func makeLogger(cfg *config.Config) (*zap.Logger, error) {
	ll := zap.InfoLevel

	switch cfg.Log.Level {
	case "debug":
		ll = zap.DebugLevel
	case "info":
		ll = zap.InfoLevel
	case "warning":
		ll = zap.WarnLevel
	case "error":
		ll = zap.ErrorLevel
	}

	_, err := os.Stat(cfg.Log.File)
	if os.IsNotExist(err) {
		dir := filepath.Dir(cfg.Log.File)

		if _, err := os.Stat(cfg.Log.File); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return nil, err
			}
		}

		file, err := os.Create(cfg.Log.File)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	zCfg := zap.NewProductionConfig()
	zCfg.Level = zap.NewAtomicLevelAt(ll)
	zCfg.OutputPaths = []string{
		cfg.Log.File,
	}

	return zCfg.Build()
}

func (a *App) Run() {
	hello := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, I am calendar!")
	}

	ver := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(
			w,
			"%s-%s-%s-%s-%s",
			version.REPO,
			version.RELEASE,
			version.BRANCH,
			version.COMMIT,
			version.DATE,
		)
	}

	middleware := func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			a.logger.Debug(
				"request received",
				zap.String("remoteAddr", r.RemoteAddr),
				zap.String("method", r.Method),
				zap.Any("URL", r.URL),
			)
			handler.ServeHTTP(w, r)
		})
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/version", ver)

	fmt.Printf("listening on %v:%v\n", a.cfg.Server.IP, a.cfg.Server.Port)
	fmt.Printf("logging to %v\n", a.cfg.Log.File)

	err := http.ListenAndServe(
		fmt.Sprintf("%s:%d", a.cfg.Server.IP, a.cfg.Server.Port),
		middleware(http.DefaultServeMux),
	)
	if err != nil {
		a.logger.Fatal("unexpected error in ListenAndServe", zap.Error(err))
	}
}
