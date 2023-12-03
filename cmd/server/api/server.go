package api

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"github.com/zcubbs/mrelay/cmd/server/config"
	"github.com/zcubbs/mrelay/cmd/server/logging"
	"github.com/zcubbs/mrelay/cmd/server/mail"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Server provides an http.Server.
type Server struct {
	*http.Server
}

type Options struct {
	Config *config.Configuration
	DbConn *bun.DB
	Mailer mail.Mailer
	logger logging.StructuredLogger

	Version string
	Commit  string
	Date    string
}

// NewServer creates and configures an APIServer serving all application routes.
func NewServer(options Options) (*Server, error) {
	log.Println("configuring server...")

	var addr string
	port := options.Config.HttpServer.Port

	if strings.Contains(port, ":") {
		addr = port
	} else {
		addr = ":" + port
	}

	api, err := New(options)
	if err != nil {
		return nil, err
	}

	srv := http.Server{
		Addr:         addr,
		Handler:      api,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return &Server{&srv}, nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *Server) Start() {
	log.Println("starting server...")
	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.Addr)

	// handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Server gracefully stopped")
}
