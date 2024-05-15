package api

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/Brix101/nestfile/internal/domain"
	"github.com/Brix101/nestfile/internal/files"
	"github.com/Brix101/nestfile/internal/middlewares"
	"github.com/Brix101/nestfile/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
)

type api struct {
	logger     *zap.Logger
	httpClient *http.Client
	assetsFs   fs.FS

	userRepo domain.UserRepository
}

func NewHTTPHandler(ctx context.Context, logger *zap.Logger, db *sql.DB, assetsFs fs.FS, fileServer *files.FileReader) *api {

	userRepo := repository.NewSqlUser(db)

	client := &http.Client{}

	return &api{
		logger:     logger,
		httpClient: client,
		assetsFs:   assetsFs,

		userRepo: userRepo,
	}
}

func (a *api) Server(port int) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: a.Routes(),
	}
}

func (a *api) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middlewares.AuthMiddleware)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://192.168.254.152/:5173", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	index, static := getStaticHandler(a.assetsFs, a.logger)

	r.Get("/vite*", static)
	r.Get("/assets/*", static)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/auth", a.authRoutes())
		r.Mount("/users", a.userRoutes())
		r.Mount("/resources", a.resourceRoutes())
	})

	r.NotFound(index)

	return r
}
