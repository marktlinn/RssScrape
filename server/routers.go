package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type ChiRouter struct {
	*chi.Mux
}

// newRouter defines a ChiMux NewRouter and enables cors middleware on the Router.
func newRouter() *ChiRouter {
	router := &ChiRouter{chi.NewRouter()}

	// temporary liberal values set for init setup.
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders:   []string{"Linke"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := &ChiRouter{chi.NewRouter()}
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)

	router.Mount("/v1", v1Router)
	return router
}
