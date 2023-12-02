package main

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/zcubbs/mrelay/cmd/server/config"
	"github.com/zcubbs/mrelay/cmd/server/db"
	"github.com/zcubbs/mrelay/cmd/server/handler"
	"github.com/zcubbs/mrelay/cmd/server/repository"
	"github.com/zcubbs/mrelay/cmd/server/web"
	"log"
	"net/http"
	"time"
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

	// create repositories
	mailRepo := repository.NewMailRepository(conn, cfg.Database)

	router := chi.NewRouter()

	// Set up middleware
	router.Use(middleware.Logger)
	// Cors middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // for development
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Create handlers and routes
	mailHandler := handler.NewMailHandler(mailRepo, cfg.Smtp, cfg.AwsSes)
	router.Mount("/api/mail", mailHandler.Routes())

	// ops handler
	opsHandler := handler.NewOpsHandler(Version, Commit, Date)
	router.Mount("/api", opsHandler.Routes())

	// Create a web app router
	router.Handle("/*", web.SPAHandler())

	// Start the server
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.HttpServer.Port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Printf("Server started on port %s\n", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
