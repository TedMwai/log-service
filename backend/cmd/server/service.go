package main

import (
	"context"
	"fmt"
	"log-management/config"
	"log-management/dao"
	"log-management/handler"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

type Service struct {
	PG     *bun.DB
	Server *http.Server
}

func New() *Service {
	return &Service{}
}

func (s *Service) Start() {
	// Initialise database connection
	pg := dao.InitPG()
	db := dao.NewPGDAO(pg)
	handler := handler.New(
		db,
		pg,
	)

	mux := initRouter(handler)

	server := &http.Server{Addr: ":" + config.AppPort, Handler: mux}

	s.Server = server
	s.PG = pg

	go func() {
		fmt.Println("Server starting on port", config.AppPort, "...")

		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to start server")
		}
	}()
}

func (s *Service) Shutdown() {
	defer os.Exit(0)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	// Block until shutfown signal received
	sig := <-sigCh

	start := time.Now()
	fmt.Printf("Received shutdown signal (%s), terminating ...\n", sig.String())

	if s.PG != nil {
		fmt.Println("Terminating database connection ...")

		if err := s.PG.Close(); err != nil {
			log.Err(err).Msg("error terminating database connection")
		}
	}

	if s.Server != nil {
		fmt.Println("Terminating http server ...")

		// Give server 10s seconds to shutdown, and complete in-flight requests
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := s.Server.Shutdown(ctx); err != nil {
			log.Err(err).Msg("error terminating http server")
		}
	}

	fmt.Printf("Shutdown completed in %fs\n", time.Since(start).Seconds())

}