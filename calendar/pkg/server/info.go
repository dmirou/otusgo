package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dmirou/otusgo/calendar/pkg/config"
	"github.com/dmirou/otusgo/calendar/pkg/version"
	"go.uber.org/zap"
)

type InfoServer struct {
	server *http.Server
	cfg    *config.Info
	logger *zap.Logger
}

func NewInfoServer(cfg *config.Info, logger *zap.Logger) *InfoServer {
	return &InfoServer{
		cfg:    cfg,
		logger: logger,
	}
}

func (is *InfoServer) Run() error {
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

	logging := func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			is.logger.Debug(
				"request received",
				zap.String("remoteAddr", r.RemoteAddr),
				zap.String("method", r.Method),
				zap.Any("URL", r.URL),
			)
			next.ServeHTTP(w, r)
		})
	}

	router := http.NewServeMux()
	// Register your routes
	router.HandleFunc("/version", logging(ver))

	addr := fmt.Sprintf("%s:%d", is.cfg.IP, is.cfg.Port)

	fmt.Printf("info server listens on %s\n", addr)

	is.server = &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return is.server.ListenAndServe()
}

func (is *InfoServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	is.server.SetKeepAlivesEnabled(false)

	if err := is.server.Shutdown(ctx); err != nil {
		is.logger.Fatal("could not gracefully shutdown the info server: %v", zap.Error(err))
		return
	}

	is.logger.Info("info server was gracefully shutdown")
}
