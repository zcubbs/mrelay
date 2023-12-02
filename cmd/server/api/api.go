package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/zcubbs/mrelay/cmd/server/api/handler"
	"github.com/zcubbs/mrelay/cmd/server/auth"
	"github.com/zcubbs/mrelay/cmd/server/db"
	"github.com/zcubbs/mrelay/cmd/server/logging"
	"github.com/zcubbs/mrelay/cmd/server/web"
	"time"
)

func New(options Options) (*chi.Mux, error) {

	//authStore := db.NewAuthStore(options.DbConn, options.Config.Database)
	authStore := db.NewInMemoryAuthStore(options.Config.Accounts)
	mailStore := db.NewMailStore(options.DbConn, options.Config.Database)

	mailHandler := handler.NewMailHandler(mailStore, options.Mailer)
	adminHandler := handler.NewAdminHandler()
	opsHandler := handler.NewOpsHandler(options.Version, options.Commit, options.Date)
	authHandler := auth.NewAuthHandler()

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Use(logging.NewStructuredLogger(options.logger.Logger))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// use CORS middleware if client is not served by this api, e.g. from other domain or CDN
	if options.Config.HttpServer.EnableCORS {
		if options.Config.HttpServer.AllowOrigins == nil {
			options.Config.HttpServer.AllowOrigins = []string{"*"}
		}
		r.Use(corsConfig(options.Config.HttpServer.AllowOrigins).Handler)
	}

	r.Mount("/auth", authHandler.Routes())
	r.Mount("/api/ops", opsHandler.Routes())
	r.Mount("/api/admin", adminHandler.Routes())
	r.Mount("/api/mail", mailHandler.Routes("accounts", authStore))

	r.Handle("/*", web.SPAHandler())

	return r, nil
}
