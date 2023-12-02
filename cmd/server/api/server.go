package api

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"github.com/zcubbs/mrelay/cmd/server/config"
	"github.com/zcubbs/mrelay/cmd/server/email"
	"github.com/zcubbs/mrelay/cmd/server/logging"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

// Server provides an http.Server.
type Server struct {
	*http.Server
}

type Options struct {
	Config *config.Configuration
	DbConn *bun.DB
	Mailer *email.Mailer
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
		Addr:    addr,
		Handler: api,
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

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)
	// teardown logic...

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Server gracefully stopped")
}
