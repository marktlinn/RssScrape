package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marktlinn/RssScrape/db"
)

type ServerConfig struct {
	DB  *db.Config
	srv *http.Server
}

func NewServer(port string, dbCfg *db.Config) *ServerConfig {
	router := newRouter()

	server := ServerConfig{
		DB: dbCfg,
		srv: &http.Server{
			Handler: router,
			Addr:    fmt.Sprintf(":%s", port),
		},
	}

	v1Router := router
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/user", server.handleCreateUser)

	router.Mount("/v1", v1Router)

	return &server
}

func (s *ServerConfig) RunServer() {
	if s.srv.Handler == nil || s.srv.Addr == "" {
		log.Fatalf("failed to run server: Handler and Addr must be specified on ServerConfig instance")
	}

	log.Printf("Server running on port: %s\n", s.srv.Addr)

	if err := s.srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to listen on port %s\n", err)
	}
}
