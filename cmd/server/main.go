package main

import (
	"flag"
	"github.com/zcubbs/mrelay/cmd/server/api"
	"github.com/zcubbs/mrelay/cmd/server/config"
	"github.com/zcubbs/mrelay/cmd/server/db"
	"github.com/zcubbs/mrelay/cmd/server/logging"
	"github.com/zcubbs/mrelay/cmd/server/mail"
	"log"
)

var (
	configPath = flag.String("config", ".", "Path to the configuration file")
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

func main() {
	flag.Parse()

	cfg, err := config.LoadConfiguration(*configPath)
	if err != nil {
		log.Fatal("Error loading configuration", "error", err)
	}

	// init db
	conn, err := db.InitializeDB(cfg.Database)
	if err != nil {
		log.Fatal("Error initializing database", "error", err)
	}

	// initialize logger
	logger := logging.NewLogger(cfg.Logging)

	// initialize mailer
	mailer := mail.NewDefaultMailer(cfg.Smtp)

	// initialize server
	srv, err := api.NewServer(api.Options{
		Config:  &cfg,
		DbConn:  conn,
		Mailer:  mailer,
		Version: Version,
		Commit:  Commit,
		Date:    Date,
	})
	if err != nil {
		logger.Fatal("Error initializing server", "error", err)
	}

	// start server
	srv.Start()
}
