package api

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
)

type api struct {
	logger     *zap.Logger
	httpClient *http.Client
	hFS        http.Handler
}

func NewHTTPHandler(ctx context.Context, logger *zap.Logger, db *sql.DB, assetsFs fs.FS) *api {

	client := &http.Client{}

	hFS := http.FileServer(http.FS(assetsFs))

	return &api{
		logger:     logger,
		httpClient: client,
		hFS:        hFS,
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

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://192.168.254.152/:5173", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "/api") {
				h.ServeHTTP(w, r)
				return
			}
			a.hFS.ServeHTTP(w, r)
		})
	})

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		a.hFS.ServeHTTP(w, r)
	})
	// r.Get("/*", http.StripPrefix("/", a.hFS).ServeHTTP)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/auth", a.AuthRoutes())
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		http.Redirect(w, r, "/index.html", http.StatusFound)
	})

	return r
}
