// Package app содержит структуру и методы для запуска приложения.
package app

import (
	"context"
	"net/http"
	"time"
)

// App представляет HTTP-сервер приложения.
type App struct{ server *http.Server }

// Run запускает HTTP-сервер на указанном хосте и порту.
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

// Shutdown корректно останавливает HTTP-сервер.
func (s *App) Shutdown(ctx context.Context) error { return s.server.Shutdown(ctx) }
