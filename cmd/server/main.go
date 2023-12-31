package main

import (
	"flag"
	"github.com/zcubbs/log"
	"github.com/zcubbs/log/structuredlogger"
	"github.com/zcubbs/mrelay/cmd/server/api"
	"github.com/zcubbs/mrelay/cmd/server/config"
	"github.com/zcubbs/mrelay/cmd/server/db"
	"github.com/zcubbs/mrelay/cmd/server/mail"
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

var cfg *config.Configuration

var configPath = flag.String("config", "", "Path to the configuration file")

func init() {
	flag.Parse()

	// initialize logger
	log.SetLogger(log.NewLogger(
		structuredlogger.CharmLoggerType,
		structuredlogger.JSONFormat,
		structuredlogger.InfoLevel,
	))
	var err error
	cfg, err = config.Load(*configPath)
	if err != nil {
		log.Fatal("failed to load configuration", "error", err)
	}

	// Load configuration
	log.Info("loaded configuration")

	cfg.BuildInfo.Version = Version
	cfg.BuildInfo.Commit = Commit
	cfg.BuildInfo.Date = Date

	log.Info("build info", "version",
		cfg.BuildInfo.Version, "commit", cfg.BuildInfo.Commit, "date", cfg.BuildInfo.Date)

	if cfg.Debug {
		log.SetLevel(structuredlogger.DebugLevel)
		config.PrintConfiguration(*cfg)
	}

	if cfg.DevMode {
		log.SetFormat(structuredlogger.TextFormat)
	} else {
		log.SetFormat(structuredlogger.JSONFormat)
	}

	log.Info("loaded configuration")
}

func main() {
	// init db
	conn, err := db.InitializeDB(cfg.Database)
	if err != nil {
		log.Fatal("Error initializing database", "error", err)
	}

	// initialize mailer
	mailer := mail.NewDefaultMailer(cfg.Smtp)

	// initialize server
	srv, err := api.NewServer(api.Options{
		Config:  cfg,
		DbConn:  conn,
		Mailer:  mailer,
		Version: Version,
		Commit:  Commit,
		Date:    Date,
	})
	if err != nil {
		log.Fatal("Error initializing server", "error", err)
	}

	// start server
	srv.Start()
}
