package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/dmirou/otusgo/calendar/pkg/config"
	"github.com/dmirou/otusgo/calendar/pkg/event"
	"github.com/dmirou/otusgo/calendar/pkg/event/repository/localcache"
	"github.com/dmirou/otusgo/calendar/pkg/event/usecase"
	"github.com/dmirou/otusgo/calendar/pkg/server"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
	is := server.NewInfoServer(a.cfg, a.logger)

	// Run info server
	go func() {
		if err := is.Run(); err != nil {
			// Check for known errors
			if err != http.ErrServerClosed {
				a.logger.Fatal(err.Error())
			}
			// Normal shutdown
			a.logger.Info(err.Error())
		}
	}()

	cs := server.NewCoreServer(a.euc, a.cfg, a.logger)

	// Run core server
	go func() {
		if err := cs.Run(); err != nil {
			// Check for known errors
			if err != grpc.ErrServerStopped {
				a.logger.Fatal(err.Error())
			}
			// Normal shutdown
			a.logger.Info(err.Error())
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	a.logger.Info("terminate signal received")

	is.Shutdown()
	cs.Shutdown()

	a.logger.Info("application finished")

	if err := a.logger.Sync(); err != nil {
		log.Fatalf("unexpected error in logger.Sync: %v", err)
	}
}
