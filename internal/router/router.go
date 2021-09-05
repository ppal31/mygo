package router

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/ppal31/mygo/internal/api/resources"
	"github.com/ppal31/mygo/internal/config"
	"github.com/ppal31/mygo/internal/store/database"
	"net/http"
)

func New(config *config.AppConfig, db *database.DataStore) http.Handler {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Use(middleware.Recoverer)
		ncors := cors.New(
			cors.Options{
				AllowedOrigins:   config.Cors.AllowedOrigins,
				AllowedMethods:   config.Cors.AllowedMethods,
				AllowedHeaders:   config.Cors.AllowedHeaders,
				ExposedHeaders:   config.Cors.ExposedHeaders,
				AllowCredentials: config.Cors.AllowCredentials,
				MaxAge:           config.Cors.MaxAge,
			},
		)
		r.Use(ncors.Handler)

		// policy endpoints
		r.Route("/resources", func(r chi.Router) {
			r.Get("/", resources.HandleList(db.ResourceStore))
			r.Get("/{resourceId}", resources.HandleGet(db.ResourceStore))
		})
	})
	return r
}
