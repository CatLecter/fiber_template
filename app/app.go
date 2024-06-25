package app

import (
	"context"
	"net/http"
	"time"
)

type App struct{ server *http.Server }

func (s *App) Run(host string, port string, handler http.Handler) error {
	s.server = &http.Server{
		Addr:           host + ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.server.ListenAndServe()
}

func (s *App) Shutdown(ctx context.Context) error { return s.server.Shutdown(ctx) }
