package main

import (
	"log-management/handler"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

func initRouter(h *handler.Handler) *chi.Mux {
	router := chi.NewRouter()

	// middleware stack
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Use(hlog.NewHandler(log.Logger))
	router.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("incoming request")
	}))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	// routes
	router.Get("/", h.Ping)
	router.Get("/ping", h.Ping)

	return router
}